package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ColleBoll/LabControl/pkg/protocol"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func New(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

/*
This function will sent the given data as agent report model to the Server at the baseURL
*/
func (c *Client) SendReport(report protocol.AgentReport) error {
	data, err := json.Marshal(report)
	if err != nil {
		return fmt.Errorf("marshal report: %w", err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.BaseURL+"/api/v1/agents/report",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("send report: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("server returned status %d", resp.StatusCode)
	}

	return nil
}
