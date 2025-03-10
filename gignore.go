package gignore

import (
	"io"

	log "github.com/sirupsen/logrus"

	"github.com/onyx-and-iris/gignore/internal/filewriter"
	"github.com/onyx-and-iris/gignore/internal/registry"
)

//go:generate go run cmd/gen/gen.go

type GignoreClient struct {
	registry *registry.TemplateRegistry
	writer   io.Writer
}

func New(options ...Option) *GignoreClient {
	gc := &GignoreClient{
		registry.New(),
		filewriter.New()}

	for _, option := range options {
		option(gc)
	}

	return gc
}

func (g *GignoreClient) List() ([]string, error) {
	return g.registry.ListTemplates()
}

func (g *GignoreClient) Create(template string) error {
	ok, err := g.registry.Contains(template)
	if err != nil {
		return err
	}
	if !ok {
		templateNotFoundErr := &templateNotFoundError{template, g.registry}
		if g.registry.Directory == "gitignoreio" {
			return templateNotFoundErr
		}

		log.Errorf("%s. Checking default registry...", templateNotFoundErr)

		g.registry.Directory = "gitignoreio"
		ok, err = g.registry.Contains(template)
		if err != nil {
			return err
		}
		if !ok {
			return templateNotFoundErr
		}
		log.Infof("template '%s' found in default gitignoreio registry", template)
	}

	content, err := g.registry.GetTemplate(template)
	if err != nil {
		return err
	}

	_, err = g.writer.Write(content)
	if err != nil {
		return err
	}
	return nil
}
