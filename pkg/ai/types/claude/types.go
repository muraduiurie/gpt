package claude

import "encoding/json"

type (
	ClaudeAIModel string
	ClaudeAIRole  string
)

const (
	// roles
	ClaudeAIModelSonnet4_20250514 ClaudeAIModel = "claude-sonnet-4-20250514"

	// models
	// roles
	ClaudeAIRoleUser      ClaudeAIRole = "user"
	ClaudeAIRoleAssistant ClaudeAIRole = "assistant"
	ClaudeAIRoleSystem    ClaudeAIRole = "system"
)

func (t *TextInputRequest) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

type TextInputRequest struct {
	Model     ClaudeAIModel             `json:"model"`
	MaxTokens int                       `json:"max_tokens,omitempty"`
	Messages  []TextInputRequestMessage `json:"messages"`
}

type TextInputRequestMessage struct {
	Role    ClaudeAIRole `json:"role"`
	Content string       `json:"content"`
}

func (t *TextInputResponse) Unmarshal(b []byte) error {
	return json.Unmarshal(b, t)
}

type TextInputResponse struct {
	Id           string                     `json:"id"`
	Type         string                     `json:"type"`
	Role         ClaudeAIRole               `json:"role"`
	Model        ClaudeAIModel              `json:"model"`
	Content      []TextInputResponseContent `json:"content"`
	StopReason   string                     `json:"stop_reason"`
	StopSequence interface{}                `json:"stop_sequence"`
	Usage        TextInputResponseUsage     `json:"usage"`
}

type TextInputResponseContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type TextInputResponseUsage struct {
	InputTokens              int                                 `json:"input_tokens"`
	CacheCreationInputTokens int                                 `json:"cache_creation_input_tokens"`
	CacheReadInputTokens     int                                 `json:"cache_read_input_tokens"`
	CacheCreation            TextInputResponseUsageCacheCreation `json:"cache_creation"`
	OutputTokens             int                                 `json:"output_tokens"`
	ServiceTier              string                              `json:"service_tier"`
}
type TextInputResponseUsageCacheCreation struct {
	Ephemeral5MInputTokens int `json:"ephemeral_5m_input_tokens"`
	Ephemeral1HInputTokens int `json:"ephemeral_1h_input_tokens"`
}
