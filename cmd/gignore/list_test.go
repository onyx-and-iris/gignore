package main

import (
	"bytes"
	"context"
	"testing"

	"github.com/onyx-and-iris/gignore"
)

func TestListSuccess(t *testing.T) {
	client := gignore.New(
		gignore.WithTemplateDirectory(gignore.DefaultTemplateDirectory),
	)
	ctx := context.WithValue(context.Background(), clientKey, client)

	testCases := []struct {
		language string
		expected string
	}{
		{
			language: "python",
			expected: "templates/gitignoreio/circuitpython.gitignore\n" +
				"templates/gitignoreio/python.gitignore\n" +
				"templates/gitignoreio/pythonvanilla.gitignore\n",
		},
	}

	var out bytes.Buffer

	// Call the listTemplates function with the context and output buffer
	err := listTemplates(ctx, &out, "python")
	if err != nil {
		t.Fatalf("listTemplates failed: %v", err)
	}

	// Check if the output contains the expected templates
	for _, tc := range testCases {
		if out.String() != tc.expected {
			t.Errorf("Expected %q, got %q", tc.expected, out.String())
		}
	}
}
