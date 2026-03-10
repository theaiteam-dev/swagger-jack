package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// DefaultBaseURL is the default API base URL embedded from the OpenAPI spec.
const DefaultBaseURL = ""

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
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
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

	// FIXME(milestone-2): MVP only supports Bearer token auth via env var.
	// API key (custom header) and Basic auth will be added in Milestone 2,
	// driven by OpenAPI securitySchemes in the parsed spec.
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
		bodyStr := string(responseBody)
		if len(bodyStr) > 200 {
			bodyStr = bodyStr[:200] + "... (truncated)"
		}
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, bodyStr)
	}

	return responseBody, nil
}
