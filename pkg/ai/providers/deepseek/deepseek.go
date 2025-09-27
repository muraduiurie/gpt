package deepseek

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	dstypes "github.com/muraduiurie/gpt/pkg/ai/types/deepseek"
	"github.com/muraduiurie/gpt/pkg/ai/types/union"
)

type Client struct {
	ApiToken          string
	TextInputEndpoint string
}

// AskAI sends a text request to the configured ChatGPT endpoint and returns
// the parsed response. It validates inputs, performs the HTTP POST request,
// and unmarshals the response body. An error is returned for invalid input,
// network issues, or unexpected HTTP status codes.
func (c *Client) AskAI(opts *union.Request) (*union.Response, error) {
	if c.TextInputEndpoint == "" {
		c.TextInputEndpoint = "https://api.deepseek.com/chat/completions"
	}
	if opts == nil {
		return nil, errors.New("nil opts")
	}
	textRequest, ok := opts.TextRequest.(*dstypes.TextInputRequest)
	if !ok {
		return nil, fmt.Errorf("*dstypes.TextInputRequest type conversion failed")
	}

	if textRequest.Model == "" {
		textRequest.Model = dstypes.DeepSeekAIModelChat
	}
	if len(textRequest.Messages) == 0 {
		return nil, fmt.Errorf("messages is required")
	}
	for i, m := range textRequest.Messages {
		if m.Role == "" {
			textRequest.Messages[i].Role = dstypes.DeepSeekAIRoleUser
		}
		if m.Content == "" {
			return nil, fmt.Errorf("content in message is required")
		}
	}

	body, err := opts.TextRequest.Marshal()
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.TextInputEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.ApiToken)

	httpClient := &http.Client{Timeout: 30 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(respBody))
	}

	var textResponse dstypes.TextInputResponse
	err = textResponse.Unmarshal(respBody)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &union.Response{
		TextResponse: &textResponse,
	}, nil
}
