// Package filewriter provides functionality to write content to a .gitignore file.
package filewriter

import (
	"bytes"
	"io"
	"os"
)

// FileWriter provides functionality to write content to a .gitignore file.
type FileWriter struct {
	targetFileName string
}

// New creates a new FileWriter with the default target file name.
func New() *FileWriter {
	return &FileWriter{".gitignore"}
}

func (fw *FileWriter) writeContent(content []byte, dst io.Writer) (int64, error) {
	r := bytes.NewReader(content)

	n, err := io.Copy(dst, r)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func (fw *FileWriter) Write(content []byte) (int, error) {
	f, err := os.Create(fw.targetFileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	const header = `# Auto-generated .gitignore by gignore: github.com/onyx-and-iris/gignore`
	const footer = `# End of gignore: github.com/onyx-and-iris/gignore`

	var sz int64
	n, err := fw.writeContent([]byte(header), f)
	if err != nil {
		return 0, err
	}
	sz += n
	n, err = fw.writeContent(content, f)
	if err != nil {
		return 0, err
	}
	sz += n
	n, err = fw.writeContent([]byte(footer), f)
	if err != nil {
		return 0, err
	}
	return int(sz + n), nil
}
