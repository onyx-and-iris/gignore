// Package main provides the entry point for the gignore command-line tool,
// which generates .gitignore files based on specified templates.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/onyx-and-iris/gignore"
	log "github.com/sirupsen/logrus"
)

func exit(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}

func main() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()

		fmt.Fprint(w, "Usage of gignore:\n")
		fmt.Fprint(w, "  gignore [flags] <template>\n")
		fmt.Fprint(w, "\n")

		fmt.Fprint(w, "Flags:\n")
		flag.PrintDefaults()

		fmt.Fprint(w, "\n")
		fmt.Fprint(w, "Example:\n")
		fmt.Fprint(w, "  gignore go\n")
	}

	var (
		list        bool
		templateDir string
		loglevel    string
	)

	flag.BoolVar(&list, "list", false, "list available templates")
	flag.BoolVar(&list, "ls", false, "list available templates (shorthand)")
	flag.StringVar(
		&templateDir,
		"dir",
		getEnv("GIGNORE_TEMPLATE_DIR", "gitignoreio"),
		"directory containing .gitignore templates",
	)
	flag.StringVar(&loglevel, "loglevel", "warn", "log level")
	flag.StringVar(&loglevel, "l", "warn", "log level (shorthand)")

	flag.Parse()

	level, err := log.ParseLevel(loglevel)
	if err != nil {
		exit(fmt.Errorf("invalid log level: %s", loglevel))
	}
	log.SetLevel(level)

	client := gignore.New(gignore.WithTemplateDirectory(templateDir))

	if list {
		if err := listTemplates(client); err != nil {
			exit(fmt.Errorf("failed to list templates: %v", err))
		}
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return
	}

	for _, arg := range args {
		err := client.Create(arg)
		if err != nil {
			exit(fmt.Errorf("failed to create .gitignore file: %v", err))
		}
		fmt.Printf("√ created %s .gitignore file\n", arg)
	}
}

func listTemplates(client *gignore.Client) error {
	templates, err := client.List()
	if err != nil {
		return err
	}
	for _, template := range templates {
		fmt.Println(template)
	}
	return nil
}
