// Package main provides the entry point for the gignore CLI tool,
// including commands like listing available .gitignore templates.
package main

import (
	"os"

	"github.com/onyx-and-iris/gignore"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var client *gignore.Client

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
		loglevel, err := log.ParseLevel(cmd.Flag("loglevel").Value.String())
		cobra.CheckErr(err)
		log.SetLevel(loglevel)

		// Initialise the gignore client
		client = gignore.New(
			gignore.WithTemplateDirectory(cmd.Flag("root").Value.String()),
		)
	},
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
	},
}

// init initialises the root command and adds global flags.
func init() {
	getEnv := func(key, defaultValue string) string {
		value := os.Getenv(key)
		if value == "" {
			return defaultValue
		}
		return value
	}

	rootCmd.PersistentFlags().
		StringP("root", "r", getEnv("GIGNORE_TEMPLATE_ROOT", gignore.DefaultTemplateDirectory), "Root directory to search for .gitignore files")
	rootCmd.PersistentFlags().
		StringP("loglevel", "l", getEnv("GIGNORE_LOGLEVEL", "warn"), "Log level (trace, debug, info, warn, error, fatal, panic)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
