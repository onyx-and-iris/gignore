// Package main generates gitignore.io templates using the gogi library.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cuonglm/gogi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Generating gitignore.io templates...")

	gogiClient, _ := gogi.NewHTTPClient()

	templates, err := fetchTemplates(gogiClient)
	if err != nil {
		log.Fatal(err)
	}

	errChan := make(chan error)

	for _, template := range templates {
		go func() {
			err := createTemplate(template)
			if err != nil {
				errChan <- fmt.Errorf("Failed to create template %s: %v", template.Name, err)
				return
			}
			errChan <- nil
		}()
	}

	for range templates {
		if err := <-errChan; err != nil {
			log.Error(err)
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
