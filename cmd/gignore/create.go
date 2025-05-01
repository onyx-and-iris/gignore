// Package main provides the entry point for the gignore CLI tool,
// including commands like listing available .gitignore templates.
package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// createCmd is the command to create a new .gitignore file.
var createCmd = &cobra.Command{
	Use:   "create",
	Args:  cobra.MinimumNArgs(1),
	Short: "Create a new .gitignore file",
	Long: `Create a new .gitignore file in the current directory.
At least one template must be specified.
Multiple templates can be specified, and they will be combined into a single .gitignore file.
Example:
	gignore create python
	gignore create python go`,
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, arg := range args {
			err := createTemplate(cmd.Context(), os.Stdout, arg)
			if err != nil {
				return fmt.Errorf("failed to create .gitignore file: %w", err)
			}
		}
		return nil
	},
}

// init initialises the create command and adds it to the root command.
func init() {
	rootCmd.AddCommand(createCmd)
}

// createTemplate creates a new .gitignore file using the specified template.
func createTemplate(ctx context.Context, out io.Writer, template string) error {
	client := getClientFromContext(ctx)
	err := client.Create(template)
	if err != nil {
		return err
	}

	fmt.Fprintf(out, "√ created %s .gitignore file\n", template)
	return nil
}
