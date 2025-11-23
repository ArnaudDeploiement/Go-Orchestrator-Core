package LLM

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OllamaClient struct {
	URL    string
	client *http.Client
}

func NewOllamaClient() *OllamaClient {
	return &OllamaClient{
		URL:    "http://localhost:11434/api/chat",
		client: &http.Client{},
	}
}

func (c *OllamaClient) Chat(req ChatRequest) (*ChatResponse, error) {
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	r, err := http.NewRequest(http.MethodPost, c.URL, bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	r.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("llm returned status %d: %s", resp.StatusCode, string(errBody))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	out := &ChatResponse{}
	err = json.Unmarshal(body, out)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return out, nil
}
