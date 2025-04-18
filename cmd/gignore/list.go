// Package main provides the entry point for the gignore CLI tool,
// including commands like listing available .gitignore templates.
package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd is the command to list all .gitignore templates.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all .gitignore files in the root template repository",
	Long: `List all .gitignore files in the root template repository.
This command will search the root template repository for .gitignore files and print their paths to the console.
The root template repository can be specified using the --root flag.
You can use this command to quickly find all available .gitignore templates.
Example:
  gignore --root=<path> list`,
	Run: func(cmd *cobra.Command, _ []string) {
		err := listTemplates(cmd.Context())
		cobra.CheckErr(err)
	},
}

// init initialises the list command and adds it to the root command.
func init() {
	rootCmd.AddCommand(listCmd)
}

// listTemplates retrieves and prints all .gitignore templates available from the gignore client.
func listTemplates(ctx context.Context) error {
	client := getClientFromContext(ctx)
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
