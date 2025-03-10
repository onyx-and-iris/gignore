package registry

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
)

//go:embed templates
var templates embed.FS

type TemplateRegistry struct {
	templates fs.FS
	Directory string
}

func New() *TemplateRegistry {
	return &TemplateRegistry{
		templates: templates,
	}
}

func (t *TemplateRegistry) filePath(name string) string {
	return fmt.Sprintf("templates/%s/%s.gitignore", t.Directory, name)
}

func (t *TemplateRegistry) Contains(name string) (bool, error) {
	_, err := fs.Stat(t.templates, t.filePath(name))
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (t *TemplateRegistry) GetTemplate(name string) ([]byte, error) {
	data, err := fs.ReadFile(t.templates, t.filePath(name))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (t *TemplateRegistry) ListTemplates() ([]string, error) {
	var paths []string

	err := fs.WalkDir(
		t.templates,
		fmt.Sprintf("templates/%s", t.Directory),
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !d.IsDir() {
				paths = append(paths, path)
			}

			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return paths, nil
}
