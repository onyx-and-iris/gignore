// Package main provides the entry point for the gignore CLI tool,
// including commands like listing available .gitignore templates.
package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all .gitignore files in the root template repository",
	Long: `List all .gitignore files in the root template repository.
This command will search the root template repository for .gitignore files and print their paths to the console.
The root template repository can be specified using the --root flag.
You can use this command to quickly find all available .gitignore templates.
Example:
  gignore --root=<path> list`,
	Run: func(_ *cobra.Command, _ []string) {
		if err := listTemplates(); err != nil {
			cobra.CheckErr(err)
		}
	},
}

// init initialises the list command and adds it to the root command.
func init() {
	rootCmd.AddCommand(listCmd)
}

// listTemplates retrieves and prints all .gitignore templates available from the gignore client.
// It takes a gignore.Client as a parameter and returns an error if the operation fails.
func listTemplates() error {
	templates, err := client.List()
	if err != nil {
		return err
	}
	var output strings.Builder
	for _, template := range templates {
		output.WriteString(template + "\n")
	}
	fmt.Print(output.String())
	return nil
}
