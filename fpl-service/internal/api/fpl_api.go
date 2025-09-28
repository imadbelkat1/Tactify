package fpl_api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/imadbelkat1/fpl-service/config"
)

type FplApiClient struct {
	HttpClient *http.Client
	UserAgent  string
}

func NewFplApiClient() *FplApiClient {
	return &FplApiClient{
		HttpClient: &http.Client{},
		UserAgent:  "FPL-Service-Client/1.0",
	}
}

func (c *FplApiClient) Get(ctx context.Context, endpoint string) ([]byte, error) {
	// Load base URL from config
	cfg := config.LoadConfig()
	baseURL := cfg.FplApi.BaseUrl

	url := fmt.Sprintf("%s%s", baseURL, endpoint)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	return body, nil
}

func (c *FplApiClient) GetAndUnmarshal(ctx context.Context, endpoint string, result interface{}) error {
	data, err := c.Get(ctx, endpoint)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("unmarshaling response: %w", err)
	}

	return nil
}
