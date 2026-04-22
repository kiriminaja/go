package kahttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
	Headers    map[string]string
}

// APIError represents a non-2xx HTTP response from the KiriminAja API.
type APIError struct {
	StatusCode int
	Status     string
	Body       string
}

func (e *APIError) Error() string {
	if e.Body != "" {
		return fmt.Sprintf("API error %d %s: %s", e.StatusCode, e.Status, e.Body)
	}
	return fmt.Sprintf("API error %d %s", e.StatusCode, e.Status)
}

type RequestOptions struct {
	Method  string
	Query   map[string]string
	Body    any
	Headers map[string]string
}

func (c *Client) RequestJSON(ctx context.Context, path string, opts RequestOptions) ([]byte, error) {
	baseURL := strings.TrimRight(c.BaseURL, "/")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	u, err := url.Parse(baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	if opts.Query != nil {
		q := u.Query()
		for k, v := range opts.Query {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	method := opts.Method
	if method == "" {
		method = http.MethodPost
	}

	var bodyReader io.Reader
	hasBody := opts.Body != nil && method != http.MethodGet && method != http.MethodDelete
	if hasBody {
		jsonBody, err := json.Marshal(opts.Body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	if hasBody {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Status:     http.StatusText(resp.StatusCode),
			Body:       string(data),
		}
	}

	return data, nil
}

func PostJSON[T any](ctx context.Context, c *Client, path string, body any) (*T, error) {
	return RequestJSONTyped[T](ctx, c, path, RequestOptions{
		Method: http.MethodPost,
		Body:   body,
	})
}

func GetJSON[T any](ctx context.Context, c *Client, path string) (*T, error) {
	return RequestJSONTyped[T](ctx, c, path, RequestOptions{
		Method: http.MethodGet,
	})
}

func DeleteJSON[T any](ctx context.Context, c *Client, path string) (*T, error) {
	return RequestJSONTyped[T](ctx, c, path, RequestOptions{
		Method: http.MethodDelete,
	})
}

func RequestJSONTyped[T any](ctx context.Context, c *Client, path string, opts RequestOptions) (*T, error) {
	data, err := c.RequestJSON(ctx, path, opts)
	if err != nil {
		return nil, err
	}
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}
