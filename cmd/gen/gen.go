package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cuonglm/gogi"
)

func main() {
	fmt.Println("Generating gitignore.io templates...")

	gogiClient, _ := gogi.NewHTTPClient()

	templates, err := fetchTemplates(gogiClient)
	if err != nil {
		log.Fatal(err)
	}

	for _, template := range templates {
		err := createTemplate(template)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create template %s: %v\n", template.Name, err)
		}
	}
}

func fetchTemplates(gogiClient *gogi.Client) (map[string]*gogi.ListJsonItem, error) {
	data, err := gogiClient.ListJson()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func createTemplate(template *gogi.ListJsonItem) error {
	file, err := os.Create(
		fmt.Sprintf("internal/registry/templates/gitignoreio/%s.gitignore", strings.ToLower(template.Name)),
	)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(template.Contents))
	if err != nil {
		return err
	}

	return nil
}
