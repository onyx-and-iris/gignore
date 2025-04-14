// Package main provides the entry point for the gignore CLI tool,
// including commands like listing available .gitignore templates.
package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd is the command to create a new .gitignore file.
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
			err := createTemplate(cmd.Context(), arg)
			cobra.CheckErr(err)
		}
	},
}

// init initialises the create command and adds it to the root command.
func init() {
	rootCmd.AddCommand(createCmd)
}

// createTemplate creates a new .gitignore file using the specified template.
func createTemplate(ctx context.Context, template string) error {
	client := getClientFromContext(ctx)
	err := client.Create(template)
	if err != nil {
		return err
	}

	fmt.Printf("√ created %s .gitignore file\n", template)
	return nil
}
