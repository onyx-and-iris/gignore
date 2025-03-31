// Package gignore provides a way to manage .gitignore files and templates.
package gignore

import (
	"fmt"
	"strings"
)

type templateNotFoundError struct {
	template          string
	templatesSearched []string
}

func (e *templateNotFoundError) Error() string {
	return fmt.Sprintf("template '%s' not found in %s registry", e.template, strings.Join(e.templatesSearched, ", "))
}
