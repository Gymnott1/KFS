package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is a reusable HTTP client for external API calls
type Client struct {
	BaseURL    string
	APIKey     string
	httpClient *http.Client
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		BaseURL: baseURL,
		APIKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) Get(path string, out interface{}) error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.BaseURL, path), nil)
	if err != nil {
		return err
	}
	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("external API error %d: %s", resp.StatusCode, string(body))
	}

	return json.NewDecoder(resp.Body).Decode(out)
}
