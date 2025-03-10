package gignore

import (
	"io"

	"github.com/onyx-and-iris/gignore/internal/filewriter"
	"github.com/onyx-and-iris/gignore/internal/registry"
	log "github.com/sirupsen/logrus"
)

//go:generate go run cmd/gen/gen.go

// Client is a client for managing .gitignore templates.
type Client struct {
	registry *registry.TemplateRegistry
	writer   io.Writer
}

// New creates a new Client with the provided options.
func New(options ...Option) *Client {
	c := &Client{
		registry.New(),
		filewriter.New(),
	}

	for _, option := range options {
		option(c)
	}

	return c
}

// List returns a list of available .gitignore templates.
func (c *Client) List() ([]string, error) {
	return c.registry.ListTemplates()
}

// Create generates a .gitignore file from the specified template.
func (c *Client) Create(template string) error {
	ok, err := c.registry.Contains(template)
	if err != nil {
		return err
	}
	if !ok {
		templateNotFoundErr := &templateNotFoundError{template, c.registry}
		if c.registry.Directory == "gitignoreio" {
			return templateNotFoundErr
		}

		log.Errorf("%s. Checking default registry...", templateNotFoundErr)

		c.registry.Directory = "gitignoreio"
		ok, err = c.registry.Contains(template)
		if err != nil {
			return err
		}
		if !ok {
			return templateNotFoundErr
		}
		log.Infof("template '%s' found in default gitignoreio registry", template)
	}

	content, err := c.registry.GetTemplate(template)
	if err != nil {
		return err
	}

	_, err = c.writer.Write(content)
	if err != nil {
		return err
	}
	return nil
}
