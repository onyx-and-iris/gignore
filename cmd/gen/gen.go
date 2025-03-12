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
	doneChan := make(chan struct{})

	for _, template := range templates {
		go func() {
			err := createTemplate(template)
			if err != nil {
				errChan <- fmt.Errorf("Failed to create template %s: %v", template.Name, err)
				return
			}
			doneChan <- struct{}{}
		}()
	}

	for range templates {
		select {
		case err := <-errChan:
			log.Error(err)
		case <-doneChan:
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
