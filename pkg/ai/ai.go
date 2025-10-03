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

type Model string

const (
	ModelChatGPT  Model = "chatgpt"
	ModelDeepSeek Model = "deepseek"
	ModelClaude   Model = "claude"
)

type AIOpts struct {
	ApiToken          string
	TextInputEndpoint string
}

// NewAIAgent initializes and returns an AI agent implementation based on the
// provided agent type. It reads configuration from `config.yaml` using Viper
// and currently supports the `chatgpt` agent. Returns an error if the agent
// type is unknown or required configuration (e.g., API token) is missing.
func NewAIAgent(model Model, conf *AIOpts) (AIAgent, error) {

	var token, endpoint string
	if conf != nil {
		v := viper.New()

		// base config: `config.yaml` (optional)
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		err := v.ReadInConfig()
		if err != nil {
			return nil, err
		}

		switch model {
		case ModelChatGPT:
			token = v.GetString("openai_api_token")
			endpoint = v.GetString("openai_text_input_endpoint")
		case ModelDeepSeek:
			token = v.GetString("deepseek_api_token")
			endpoint = v.GetString("deepseek_text_input_endpoint")
		case ModelClaude:
			token = v.GetString("claude_api_token")
			endpoint = v.GetString("claude_text_input_endpoint")
		default:
			return nil, fmt.Errorf("unknown ai model: %s", model)
		}
	} else {
		token = conf.ApiToken
		endpoint = conf.TextInputEndpoint
	}

	switch model {
	case ModelChatGPT:
		c := &chatgpt.Client{
			ApiToken:          token,
			TextInputEndpoint: endpoint,
		}

		if c.ApiToken == "" {
			return nil, errors.New("missing API token: set `openai_api_token` in `config.yaml`")
		}

		return c, nil
	case ModelDeepSeek:
		c := &deepseek.Client{
			ApiToken:          token,
			TextInputEndpoint: endpoint,
		}

		if c.ApiToken == "" {
			return nil, errors.New("missing API token: set `deepseek_api_token` in `config.yaml`")
		}

		return c, nil
	case ModelClaude:
		c := &claude.Client{
			ApiToken:          token,
			TextInputEndpoint: endpoint,
		}

		if c.ApiToken == "" {
			return nil, errors.New("missing API token: set `claude_api_token` in `config.yaml`")
		}

		return c, nil
	default:
		return nil, fmt.Errorf("unknown ai model: %s", model)
	}
}
