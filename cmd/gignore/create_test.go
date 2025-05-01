package main

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/onyx-and-iris/gignore"
)

// mockFileWriter is a mock implementation of the gignore.FileWriter interface.
// It writes the content to an io.Writer and adds a header and footer.
// This is used for testing purposes to avoid writing to the actual file system.
type mockFileWriter struct {
	out            io.Writer
	TargetFileName string
}

// Write implements the io.Writer interface for the mockFileWriter.
func (m mockFileWriter) Write(content []byte) (int, error) {
	const header = "# Auto-generated .gitignore by gignore: github.com/onyx-and-iris/gignore\n"
	const footer = "\n# End of gignore: github.com/onyx-and-iris/gignore\n"

	m.out.Write([]byte(header))
	m.out.Write(content)
	m.out.Write([]byte(footer))

	return len(content) + len(header) + len(footer), nil
}

func TestCreateSuccess(t *testing.T) {
	var templateBuffer bytes.Buffer
	client := gignore.New(
		gignore.WithTemplateDirectory(gignore.DefaultTemplateDirectory),
		gignore.WithFileWriter(mockFileWriter{out: &templateBuffer}),
	)
	ctx := context.WithValue(context.Background(), clientKey, client)

	var outBuffer bytes.Buffer
	err := createTemplate(ctx, &outBuffer, "go")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedContent, err := os.ReadFile("testdata/go.gitignore")
	if err != nil {
		t.Fatalf("failed to read testdata/go.gitignore: %v", err)
	}

	if templateBuffer.String() != string(expectedContent) {
		t.Fatalf("expected template content %q, got %q", string(expectedContent), templateBuffer.String())
	}

	expectedOutput := "√ created go .gitignore file\n"
	if outBuffer.String() != expectedOutput {
		t.Fatalf("expected %q, got %q", expectedOutput, outBuffer.String())
	}
}
