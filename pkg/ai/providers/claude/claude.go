package claude

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	cltypes "github.com/muraduiurie/gpt/pkg/ai/types/claude"
	"github.com/muraduiurie/gpt/pkg/ai/types/union"
)

type Client struct {
	ApiToken          string
	TextInputEndpoint string
}

func (c *Client) AskAI(opts *union.Request) (*union.Response, error) {
	if c.TextInputEndpoint == "" {
		c.TextInputEndpoint = "https://api.openai.com/v1/responses"
	}
	if opts == nil {
		return nil, errors.New("nil opts")
	}
	textRequest, ok := opts.TextRequest.(*cltypes.TextInputRequest)
	if !ok {
		return nil, fmt.Errorf("*cltypes.TextInputRequest type conversion failed")
	}

	if textRequest.Model == "" {
		textRequest.Model = cltypes.ClaudeAIModelSonnet4_20250514
	}
	if textRequest.MaxTokens == 0 {
		textRequest.MaxTokens = 100
	}
	if len(textRequest.Messages) == 0 {
		return nil, fmt.Errorf("messages is required")
	}
	for i, m := range textRequest.Messages {
		if m.Role == "" {
			textRequest.Messages[i].Role = cltypes.ClaudeAIRoleUser
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
	req.Header.Set("x-api-key", c.ApiToken)
	req.Header.Set("anthropic-version", "2023-06-01")

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

	var textResponse cltypes.TextInputResponse
	err = textResponse.Unmarshal(respBody)
	if err != nil {
		return nil, err
	}

	return &union.Response{
		TextResponse: &textResponse,
	}, nil
}
