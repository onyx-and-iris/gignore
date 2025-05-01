// Package main provides the entry point for the gignore CLI tool,
// including commands like listing available .gitignore templates.
package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/onyx-and-iris/gignore"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd is the root command for the gignore CLI tool.
var rootCmd = &cobra.Command{
	Use:   "gignore",
	Short: "A command line tool to manage .gitignore files",
	Long: `gignore is a command line tool that helps you manage your .gitignore files.
You can use it to list available templates and create new .gitignore files.
It supports various programming languages.
Example:
	gignore list
	gignore create python`,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		// Initialise the logger
		level, err := log.ParseLevel(viper.GetString("loglevel"))
		cobra.CheckErr(err)
		log.SetLevel(level)

		// Initialise the gignore client
		client := gignore.New(
			gignore.WithTemplateDirectory(viper.GetString("root")),
		)

		// Set the client in the context
		// This allows us to access the client in the command handlers
		ctx := context.WithValue(context.Background(), clientKey, client)
		cmd.SetContext(ctx)
	},
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
	},
}

var cfgFile string

// init initialises the root command and adds global flags.
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().
		StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.config/gignore/config.yaml)")
	rootCmd.PersistentFlags().
		StringP("root", "r", gignore.DefaultTemplateDirectory, "Root directory to search for .gitignore files")
	rootCmd.PersistentFlags().
		StringP("loglevel", "l", "warn", "Log level (trace, debug, info, warn, error, fatal, panic)")

	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("GIGNORE")
	viper.BindPFlag("root", rootCmd.PersistentFlags().Lookup("root"))
	viper.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))
}

// initConfig reads in config file and ENV variables if set.
// It first checks if a config file is provided via the --config flag.
// If not, it looks for a config file in the user's config directory.
// The config file is expected to be in YAML format and named "config.yaml".
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		configDir, err := os.UserConfigDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(filepath.Join(configDir, "gignore"))
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		// If the config file is not found, we ignore the error
		// as it is not mandatory to have a config file.
		// However, if there is any other error, we log it and exit.
		// This includes errors like invalid syntax in the config file.
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal("Error reading config file: ", err)
		}
	}
	viper.AutomaticEnv()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
