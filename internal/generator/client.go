// Package generator produces Go source files for the generated CLI project.
package generator

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/queso/swagger-jack/internal/model"
)

// clientTemplate is the Go source template for the generated HTTP client.
// It embeds the spec's BaseURL as DefaultBaseURL and falls back to it in NewClient.
const clientTemplate = `package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DefaultBaseURL is the default API base URL embedded from the OpenAPI spec.
const DefaultBaseURL = {{goString .BaseURL}}

// Client holds the configuration for making authenticated HTTP requests.
type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
}

// NewClient constructs a Client with the given baseURL and token.
// When baseURL is empty, DefaultBaseURL is used.
func NewClient(baseURL, token string) *Client {
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	return &Client{
		BaseURL:    baseURL,
		Token:      token,
		HTTPClient: &http.Client{},
	}
}

// Do executes an HTTP request against the API.
//
// method is the HTTP verb (GET, POST, etc.).
// path is the URL path template (e.g., "/users/{userId}").
// pathParams maps placeholder names to their runtime values for path interpolation.
// queryParams maps query parameter names to values appended to the URL.
// body is an optional request body; pass nil for requests without a body.
//
// Path parameter substitution uses strings.NewReplacer to replace {param}
// placeholders with the corresponding values from pathParams.
func (c *Client) Do(method, path string, pathParams map[string]string, queryParams map[string]string, body interface{}) ([]byte, error) {
	// Interpolate {param} placeholders in the path template.
	pairs := make([]string, 0, len(pathParams)*2)
	for key, value := range pathParams {
		pairs = append(pairs, "{"+key+"}", value)
	}
	interpolatedPath := strings.NewReplacer(pairs...).Replace(path)

	requestURL := strings.TrimRight(c.BaseURL, "/") + interpolatedPath

	// Append query parameters.
	if len(queryParams) > 0 {
		separator := "?"
		for key, value := range queryParams {
			requestURL += separator + key + "=" + url.QueryEscape(value)
			separator = "&"
		}
	}

	// Encode the request body as JSON when provided.
	var bodyReader io.Reader
	if body != nil {
		encoded, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("encoding request body: %w", err)
		}
		bodyReader = bytes.NewReader(encoded)
	}

	req, err := http.NewRequest(method, requestURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Inject Bearer token auth when a token is configured.
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	// Return a descriptive error for non-2xx responses.
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(responseBody))
	}

	return responseBody, nil
}
`

// clientTemplateData holds the values interpolated into clientTemplate.
type clientTemplateData struct {
	BaseURL string
}

// GenerateClient returns the Go source code for the generated project's HTTP
// client (internal/client.go). It embeds spec.BaseURL as DefaultBaseURL and
// generates a NewClient that falls back to it when an empty string is passed.
func GenerateClient(spec *model.APISpec) (string, error) {
	if spec == nil {
		return "", fmt.Errorf("spec must not be nil")
	}

	funcMap := template.FuncMap{
		"goString": func(s string) string {
			return fmt.Sprintf("%q", s)
		},
	}
	tmpl, err := template.New("client").Funcs(funcMap).Parse(clientTemplate)
	if err != nil {
		return "", fmt.Errorf("parsing client template: %w", err)
	}

	data := clientTemplateData{
		BaseURL: spec.BaseURL,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("executing client template: %w", err)
	}

	return strings.TrimLeft(buf.String(), "\n"), nil
}
