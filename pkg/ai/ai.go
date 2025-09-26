package ai

import (
	"errors"
	"fmt"

	"github.com/muraduiurie/gpt/pkg/ai/types/common"
	"github.com/muraduiurie/gpt/pkg/chatgpt"
	"github.com/spf13/viper"
)

type AiModel string

const (
	// chatgpt models
	AiModelGpt3_5_turbo_0301 AiModel = "gpt-3.5-turbo-0301"
	AiModelGpt3_5_turbo_1106 AiModel = "gpt-3.5-turbo-1106"
	AiModelGpt3_5_turbo      AiModel = "gpt-3.5-turbo"
	AiModelGpt4o             AiModel = "gpt-4o"
	AiModelGpt4_1            AiModel = "gpt-4.1"
	AiModelGpt4oMini         AiModel = "gpt-4o-mini"
	AiModelGpt4o_turbo       AiModel = "gpt-4-turbo"
	AiModelTTS1              AiModel = "tts-1"
	AiModelTTS1_HD           AiModel = "tts-1-hd"

	// dataseek models
	AiModelDeepSeekChat     AiModel = "deepseek-chat"
	AiModelDeepSeekReasoner AiModel = "deepseek-reasoner"
)

type AIAgent interface {
	AskAI(opts *common.Request) (*common.Response, error)
}

type Agent string

const (
	ChatGPTAgent  Agent = "chatgpt"
	DeepSeekAgent Agent = "deepseek"
)

// NewAIAgent initializes and returns an AI agent implementation based on the
// provided agent type. It reads configuration from `config.yaml` using Viper
// and currently supports the `chatgpt` agent. Returns an error if the agent
// type is unknown or required configuration (e.g., API token) is missing.
func NewAIAgent(agent Agent) (AIAgent, error) {

	v := viper.New()

	// base config: `config.yaml` (optional)
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	switch agent {
	case ChatGPTAgent:
		c := &chatgpt.Client{
			ApiToken:          v.GetString("api_token"),
			TextInputEndpoint: v.GetString("text_input_endpoint"),
		}

		if c.ApiToken == "" {
			return nil, errors.New("missing API token: set `api_token` in `config.yaml` or `OPENAI_API_KEY`")
		}

		return c, nil
	default:
		return nil, fmt.Errorf("unknown ai agent: %s", agent)
	}
}

type AiOpts struct {
	Message string
	Model   AiModel
}
