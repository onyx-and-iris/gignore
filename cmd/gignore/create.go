// Package main provides the entry point for the gignore CLI tool,
// including commands like listing available .gitignore templates.
package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new .gitignore file",
	Long: `Create a new .gitignore file in the current directory.
At least one template must be specified.
Multiple templates can be specified, and they will be combined into a single .gitignore file.
Example:
  gignore create python
  gignore create python go`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		for _, arg := range args {
			createTemplate(arg)
		}
	},
}

// init initialises the create command and adds it to the root command.
func init() {
	rootCmd.AddCommand(createCmd)
}

// createTemplate creates a new .gitignore file using the specified template.
func createTemplate(template string) {
	err := client.Create(template)
	cobra.CheckErr(err)

	fmt.Printf("√ created %s .gitignore file\n", template)
}
