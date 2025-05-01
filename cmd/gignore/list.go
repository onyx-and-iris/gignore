// Package main provides the entry point for the gignore CLI tool,
// including commands like listing available .gitignore templates.
package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd is the command to list all .gitignore templates.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all .gitignore files in the root template repository",
	Long: `List all .gitignore files in the root template repository.
This command will search the root template repository for .gitignore files and print their paths to the console.
You can also provide a pattern to filter the results.
Example:
	gignore list
	gignore list python`,
	RunE: func(cmd *cobra.Command, args []string) error {
		patterns := args
		if len(patterns) == 0 {
			patterns = []string{""}
		}
		err := listTemplates(cmd.Context(), os.Stdout, patterns...)
		if err != nil {
			return fmt.Errorf("failed to list templates: %w", err)
		}
		return nil
	},
}

// init initialises the list command and adds it to the root command.
func init() {
	rootCmd.AddCommand(listCmd)
}

// listTemplates retrieves and prints all .gitignore templates available from the gignore client.
func listTemplates(ctx context.Context, out io.Writer, patterns ...string) error {
	client := getClientFromContext(ctx)
	templates, err := client.List(patterns...)
	if err != nil {
		return err
	}

	if len(templates) == 0 {
		fmt.Println("No templates found.")
		return nil
	}

	var output strings.Builder
	for _, template := range templates {
		output.WriteString(template + "\n")
	}
	fmt.Fprint(out, output.String())

	return nil
}
