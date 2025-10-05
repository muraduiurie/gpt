package chatgpt

import (
	"encoding/json"
)

type (
	ChatGPTAIModel string
	ChatGPTAIRole  string
)

const (
	// chatgpt models
	AiModelGpt3_5_turbo_0301 ChatGPTAIModel = "gpt-3.5-turbo-0301"
	AiModelGpt3_5_turbo_1106 ChatGPTAIModel = "gpt-3.5-turbo-1106"
	AiModelGpt3_5_turbo      ChatGPTAIModel = "gpt-3.5-turbo"
	AiModelGpt4o             ChatGPTAIModel = "gpt-4o"
	AiModelGpt4_1            ChatGPTAIModel = "gpt-4.1"
	AiModelGpt4oMini         ChatGPTAIModel = "gpt-4o-mini"
	AiModelGpt4o_turbo       ChatGPTAIModel = "gpt-4-turbo"
	AiModelTTS1              ChatGPTAIModel = "tts-1"
	AiModelTTS1_HD           ChatGPTAIModel = "tts-1-hd"

	// roles
	ChatGPTAIRoleUser      ChatGPTAIRole = "user"
	ChatGPTAIRoleAssistant ChatGPTAIRole = "assistant"
	ChatGPTAIRoleSystem    ChatGPTAIRole = "system"
)

type TextInputRequest struct {
	Model ChatGPTAIModel `json:"model"`
	Input string         `json:"input"`
}

func (t *TextInputRequest) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

type ImageInputRequestContent struct {
	Type     string `json:"type"`
	Text     string `json:"text,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}

type FileInputRequestContent struct {
	Type    string `json:"type"`
	Text    string `json:"text,omitempty"`
	FileUrl string `json:"file_url,omitempty"`
}

type ImageInputRequestInput struct {
	Role    ChatGPTAIRole              `json:"role"`
	Content []ImageInputRequestContent `json:"content"`
}

type ImageInputRequest struct {
	Model ChatGPTAIModel           `json:"model"`
	Input []ImageInputRequestInput `json:"input"`
}

type FileInputRequestInput struct {
	Role    ChatGPTAIRole             `json:"role"`
	Content []FileInputRequestContent `json:"content"`
}

type FileInputRequest struct {
	Model ChatGPTAIModel          `json:"model"`
	Input []FileInputRequestInput `json:"input"`
}

// responses
type ImageInputResponseContent struct {
	Type        string        `json:"type"`
	Text        string        `json:"text"`
	Annotations []interface{} `json:"annotations"`
}

type ImageInputResponseOutput struct {
	Type    string                      `json:"type"`
	Id      string                      `json:"id"`
	Status  string                      `json:"status"`
	Role    ChatGPTAIRole               `json:"role"`
	Content []ImageInputResponseContent `json:"content"`
}

type ResponseReasoning struct {
	Effort  interface{} `json:"effort"`
	Summary interface{} `json:"summary"`
}

type ResponseMetadata struct {
}

type ImageInputResponse struct {
	Id                 string                     `json:"id"`
	Object             string                     `json:"object"`
	CreatedAt          int                        `json:"created_at"`
	Status             string                     `json:"status"`
	Error              interface{}                `json:"error"`
	IncompleteDetails  interface{}                `json:"incomplete_details"`
	Instructions       interface{}                `json:"instructions"`
	MaxOutputTokens    interface{}                `json:"max_output_tokens"`
	Model              ChatGPTAIModel             `json:"model"`
	Output             []ImageInputResponseOutput `json:"output"`
	ParallelToolCalls  bool                       `json:"parallel_tool_calls"`
	PreviousResponseId interface{}                `json:"previous_response_id"`
	Reasoning          ResponseReasoning          `json:"reasoning"`
	Store              bool                       `json:"store"`
	Temperature        float64                    `json:"temperature"`
	Text               ResponseText               `json:"text"`
	ToolChoice         string                     `json:"tool_choice"`
	Tools              []interface{}              `json:"tools"`
	TopP               float64                    `json:"top_p"`
	Truncation         string                     `json:"truncation"`
	Usage              ResponseUsage              `json:"usage"`
	User               interface{}                `json:"user"`
	Metadata           ResponseMetadata           `json:"metadata"`
}

type TextInputResponseOutputContent struct {
	Type        string        `json:"type"`
	Text        string        `json:"text"`
	Annotations []interface{} `json:"annotations"`
}

type TextInputResponseOutput struct {
	Type    string                           `json:"type"`
	Id      string                           `json:"id"`
	Status  string                           `json:"status"`
	Role    ChatGPTAIRole                    `json:"role"`
	Content []TextInputResponseOutputContent `json:"content"`
}

type ResponseTextFormat struct {
	Type string `json:"type"`
}

type ResponseText struct {
	Format ResponseTextFormat `json:"format"`
}

type ResponseUsageInputTokensDetails struct {
	CachedTokens int `json:"cached_tokens"`
}

type ResponseUsageOutputTokensDetails struct {
	ReasoningTokens int `json:"reasoning_tokens"`
}

type ResponseUsage struct {
	InputTokens         int                              `json:"input_tokens"`
	InputTokensDetails  ResponseUsageInputTokensDetails  `json:"input_tokens_details"`
	OutputTokens        int                              `json:"output_tokens"`
	OutputTokensDetails ResponseUsageOutputTokensDetails `json:"output_tokens_details"`
	TotalTokens         int                              `json:"total_tokens"`
}

func (t *TextInputResponse) Unmarshal(b []byte) error {
	return json.Unmarshal(b, t)
}

type TextInputResponse struct {
	Id                 string                    `json:"id"`
	Object             string                    `json:"object"`
	CreatedAt          int                       `json:"created_at"`
	Status             string                    `json:"status"`
	Error              interface{}               `json:"error"`
	IncompleteDetails  interface{}               `json:"incomplete_details"`
	Instructions       interface{}               `json:"instructions"`
	MaxOutputTokens    interface{}               `json:"max_output_tokens"`
	Model              ChatGPTAIModel            `json:"model"`
	Output             []TextInputResponseOutput `json:"output"`
	ParallelToolCalls  bool                      `json:"parallel_tool_calls"`
	PreviousResponseId interface{}               `json:"previous_response_id"`
	Reasoning          ResponseReasoning         `json:"reasoning"`
	Store              bool                      `json:"store"`
	Temperature        float64                   `json:"temperature"`
	Text               ResponseText              `json:"text"`
	ToolChoice         string                    `json:"tool_choice"`
	Tools              []interface{}             `json:"tools"`
	TopP               float64                   `json:"top_p"`
	Truncation         string                    `json:"truncation"`
	Usage              ResponseUsage             `json:"usage"`
	User               interface{}               `json:"user"`
	Metadata           ResponseMetadata          `json:"metadata"`
}

type FileInputResponseOutputContent struct {
	Type        string        `json:"type"`
	Annotations []interface{} `json:"annotations"`
	Logprobs    []interface{} `json:"logprobs"`
	Text        string        `json:"text"`
}

type FileInputResponseOutput struct {
	Id      string                           `json:"id"`
	Type    string                           `json:"type"`
	Status  string                           `json:"status"`
	Content []FileInputResponseOutputContent `json:"content"`
	Role    ChatGPTAIRole                    `json:"role"`
}

type FileInputResponse struct {
	Id                 string                    `json:"id"`
	Object             string                    `json:"object"`
	CreatedAt          int                       `json:"created_at"`
	Status             string                    `json:"status"`
	Background         bool                      `json:"background"`
	Error              interface{}               `json:"error"`
	IncompleteDetails  interface{}               `json:"incomplete_details"`
	Instructions       interface{}               `json:"instructions"`
	MaxOutputTokens    interface{}               `json:"max_output_tokens"`
	MaxToolCalls       interface{}               `json:"max_tool_calls"`
	Model              ChatGPTAIModel            `json:"model"`
	Output             []FileInputResponseOutput `json:"output"`
	ParallelToolCalls  bool                      `json:"parallel_tool_calls"`
	PreviousResponseId interface{}               `json:"previous_response_id"`
	Reasoning          ResponseReasoning         `json:"reasoning"`
	ServiceTier        string                    `json:"service_tier"`
	Store              bool                      `json:"store"`
	Temperature        float64                   `json:"temperature"`
	Text               ResponseText              `json:"text"`
	ToolChoice         string                    `json:"tool_choice"`
	Tools              []interface{}             `json:"tools"`
	TopLogprobs        int                       `json:"top_logprobs"`
	TopP               float64                   `json:"top_p"`
	Truncation         string                    `json:"truncation"`
	Usage              ResponseUsage             `json:"usage"`
	User               interface{}               `json:"user"`
	Metadata           ResponseMetadata          `json:"metadata"`
}
