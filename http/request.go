package kahttp

import (
	"bytes"
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

type RequestOptions struct {
	Method  string
	Query   map[string]string
	Body    any
	Headers map[string]string
}

func (c *Client) RequestJSON(path string, opts RequestOptions) ([]byte, error) {
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

	req, err := http.NewRequest(method, u.String(), bodyReader)
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
		return nil, fmt.Errorf("request failed: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return data, nil
}

func PostJSON[T any](c *Client, path string, body any) (*T, error) {
	return RequestJSONTyped[T](c, path, RequestOptions{
		Method: http.MethodPost,
		Body:   body,
	})
}

func GetJSON[T any](c *Client, path string) (*T, error) {
	return RequestJSONTyped[T](c, path, RequestOptions{
		Method: http.MethodGet,
	})
}

func DeleteJSON[T any](c *Client, path string) (*T, error) {
	return RequestJSONTyped[T](c, path, RequestOptions{
		Method: http.MethodDelete,
	})
}

func RequestJSONTyped[T any](c *Client, path string, opts RequestOptions) (*T, error) {
	data, err := c.RequestJSON(path, opts)
	if err != nil {
		return nil, err
	}
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &result, nil
}
