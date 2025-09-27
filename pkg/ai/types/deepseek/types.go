package deepseek

import "encoding/json"

type (
	DeepSeekAIModel string
	DeepSeekAIRole  string
)

const (
	// models
	DeepSeekAIModelChat     DeepSeekAIModel = "deepseek-chat"
	DeepSeekAIModelReasoner DeepSeekAIModel = "deepseek-reasoner"

	// roles
	DeepSeekAIRoleUser      DeepSeekAIRole = "user"
	DeepSeekAIRoleAssistant DeepSeekAIRole = "assistant"
	DeepSeekAIRoleSystem    DeepSeekAIRole = "system"
)

func (t *TextInputRequest) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

type TextInputRequest struct {
	Messages         []TextInputRequestMessage       `json:"messages"`
	Model            DeepSeekAIModel                 `json:"model"`
	FrequencyPenalty *int                            `json:"frequency_penalty,omitempty"`
	MaxTokens        *int                            `json:"max_tokens,omitempty"`
	PresencePenalty  *int                            `json:"presence_penalty,omitempty"`
	ResponseFormat   *TextInputRequestResponseFormat `json:"response_format,omitempty"`
	Stop             interface{}                     `json:"stop,omitempty"`
	Stream           bool                            `json:"stream,omitempty"`
	StreamOptions    interface{}                     `json:"stream_options,omitempty"`
	Temperature      *int                            `json:"temperature,omitempty"`
	TopP             *int                            `json:"top_p,omitempty"`
	Tools            interface{}                     `json:"tools,omitempty"`
	ToolChoice       *string                         `json:"tool_choice,omitempty"`
	Logprobs         bool                            `json:"logprobs,omitempty"`
	TopLogprobs      interface{}                     `json:"top_logprobs,omitempty"`
}

type TextInputRequestResponseFormat struct {
	Type string `json:"type"`
}

type TextInputRequestMessage struct {
	Content string         `json:"content"`
	Role    DeepSeekAIRole `json:"role"`
}

func (t *TextInputResponse) Unmarshal(b []byte) error {
	return json.Unmarshal(b, t)
}

type TextInputResponse struct {
	Id                string                    `json:"id"`
	Object            string                    `json:"object"`
	Created           int                       `json:"created"`
	Model             DeepSeekAIModel           `json:"model"`
	Choices           []TextInputResponseChoice `json:"choices"`
	Usage             TextInputResponseUsage    `json:"usage"`
	SystemFingerprint string                    `json:"system_fingerprint"`
}

type TextInputResponseUsage struct {
	PromptTokens          int                                       `json:"prompt_tokens"`
	CompletionTokens      int                                       `json:"completion_tokens"`
	TotalTokens           int                                       `json:"total_tokens"`
	PromptTokensDetails   TextInputResponseUsagePromptTokensDetails `json:"prompt_tokens_details"`
	PromptCacheHitTokens  int                                       `json:"prompt_cache_hit_tokens"`
	PromptCacheMissTokens int                                       `json:"prompt_cache_miss_tokens"`
}

type TextInputResponseUsagePromptTokensDetails struct {
	CachedTokens int `json:"cached_tokens"`
}

type TextInputResponseChoice struct {
	Index        int                            `json:"index"`
	Message      TextInputResponseChoiceMessage `json:"message"`
	Logprobs     interface{}                    `json:"logprobs"`
	FinishReason string                         `json:"finish_reason"`
}

type TextInputResponseChoiceMessage struct {
	Role    DeepSeekAIRole `json:"role"`
	Content string         `json:"content"`
}
