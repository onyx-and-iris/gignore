package gignore

import (
	"io"

	"github.com/onyx-and-iris/gignore/internal/filewriter"
	"github.com/onyx-and-iris/gignore/internal/registry"
	log "github.com/sirupsen/logrus"
)

//go:generate go run cmd/gen/main.go

// DefaultTemplateDirectory is the default directory for .gitignore templates.
const DefaultTemplateDirectory = "gitignoreio"

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
	return c.registry.List()
}

// Create generates a .gitignore file from the specified template.
func (c *Client) Create(template string) error {
	ok, err := c.registry.Contains(template)
	if err != nil {
		return err
	}
	if !ok {
		templateNotFoundErr := &templateNotFoundError{template, []string{c.registry.Directory}}
		if c.registry.Directory == DefaultTemplateDirectory {
			return templateNotFoundErr
		}

		c.registry.Directory = DefaultTemplateDirectory
		ok, err = c.registry.Contains(template)
		if err != nil {
			return err
		}
		if !ok {
			templateNotFoundErr.templatesSearched = append(templateNotFoundErr.templatesSearched, c.registry.Directory)
			return templateNotFoundErr
		}
		log.Debugf("template '%s' found in gitignoreio registry", template)
	} else {
		log.Debugf("template '%s' found in %s registry", template, c.registry.Directory)
	}

	content, err := c.registry.Get(template)
	if err != nil {
		return err
	}

	_, err = c.writer.Write(content)
	if err != nil {
		return err
	}
	return nil
}
