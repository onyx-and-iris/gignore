package gignore

import (
	"fmt"

	"github.com/onyx-and-iris/gignore/internal/registry"
)

type templateNotFoundError struct {
	template string
	registry *registry.TemplateRegistry
}

func (e *templateNotFoundError) Error() string {
	return fmt.Sprintf("template '%s' not found in %s registry", e.template, e.registry.Directory)
}
