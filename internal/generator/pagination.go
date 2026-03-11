package generator

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"
)

//go:embed templates/pagination.go.tmpl
var paginationTemplate string

// GeneratePagination returns the Go source code for the generated project's
// internal/client/pagination.go file. It contains the FetchAll helper and
// cursor extraction logic for all three pagination types.
func GeneratePagination(cliName string) (string, error) {
	if cliName == "" {
		return "", fmt.Errorf("cliName must not be empty")
	}

	tmpl, err := template.New("pagination").Parse(paginationTemplate)
	if err != nil {
		return "", fmt.Errorf("parsing pagination template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, struct{ CLIName string }{cliName}); err != nil {
		return "", fmt.Errorf("executing pagination template: %w", err)
	}

	return strings.TrimLeft(buf.String(), "\n"), nil
}
