package ai

import (
	"errors"
	"fmt"

	"github.com/muraduiurie/gpt/pkg/ai/providers/chatgpt"
	"github.com/muraduiurie/gpt/pkg/ai/providers/claude"
	"github.com/muraduiurie/gpt/pkg/ai/providers/deepseek"
	"github.com/muraduiurie/gpt/pkg/ai/types/union"
	"github.com/spf13/viper"
)

type AIAgent interface {
	AskAI(opts *union.Request) (*union.Response, error)
}

type Agent string

const (
	ChatGPTAgent  Agent = "chatgpt"
	DeepSeekAgent Agent = "deepseek"
	ClaudeAgent   Agent = "claude"
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
			ApiToken:          v.GetString("openai_api_token"),
			TextInputEndpoint: v.GetString("openai_text_input_endpoint"),
		}

		if c.ApiToken == "" {
			return nil, errors.New("missing API token: set `openai_api_token` in `config.yaml`")
		}

		return c, nil
	case DeepSeekAgent:
		c := &deepseek.Client{
			ApiToken:          v.GetString("deepseek_api_token"),
			TextInputEndpoint: v.GetString("deepseek_text_input_endpoint"),
		}

		if c.ApiToken == "" {
			return nil, errors.New("missing API token: set `deepseek_api_token` in `config.yaml`")
		}

		return c, nil
	case ClaudeAgent:
		c := &claude.Client{
			ApiToken:          v.GetString("claude_api_token"),
			TextInputEndpoint: v.GetString("claude_text_input_endpoint"),
		}

		if c.ApiToken == "" {
			return nil, errors.New("missing API token: set `claude_api_token` in `config.yaml`")
		}

		return c, nil
	default:
		return nil, fmt.Errorf("unknown ai agent: %s", agent)
	}
}
