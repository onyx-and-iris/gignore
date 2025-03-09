package main

import (
	"flag"
	"fmt"
	"slices"

	"github.com/onyx-and-iris/gignore"
	log "github.com/sirupsen/logrus"
)

func main() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()

		fmt.Fprint(w, "Usage of gignore:\n")
		fmt.Fprintf(w, "  gignore [flags] <template>\n")
		fmt.Fprint(w, "\n")

		fmt.Fprint(w, "Flags:\n")
		flag.PrintDefaults()

		fmt.Fprint(w, "\n")
		fmt.Fprintf(w, "Example:\n")
		fmt.Fprint(w, "  gignore go\n")
	}

	var (
		list        bool
		templateDir string
		loglevel    int
	)

	flag.BoolVar(&list, "list", false, "list available templates")
	flag.BoolVar(&list, "ls", false, "list available templates (shorthand)")
	flag.StringVar(
		&templateDir,
		"dir",
		getEnv("GIGNORE_TEMPLATE_DIR", "gitignoreio"),
		"directory containing .gitignore templates",
	)
	flag.IntVar(&loglevel, "loglevel", int(log.WarnLevel), "log level")
	flag.IntVar(&loglevel, "l", int(log.WarnLevel), "log level (shorthand)")
	flag.Parse()

	if slices.Contains(log.AllLevels, log.Level(loglevel)) {
		log.SetLevel(log.Level(loglevel))
	}

	client := gignore.New(gignore.WithTemplateDirectory(templateDir))

	if list {
		listTemplates(client)
		return
	}

	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		return
	}

	err := client.Create(args[0])
	if err != nil {
		log.Fatalf("failed to create .gitignore file: %v", err)
	}

	fmt.Printf("√ created %s .gitignore file\n", args[0])
}

func listTemplates(client *gignore.GignoreClient) {
	templates, err := client.List()
	if err != nil {
		log.Fatalf("failed to list templates: %v", err)
	}
	for _, template := range templates {
		fmt.Println(template)
	}
}
