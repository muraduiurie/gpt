## Go AI Client Library

Lightweight Go client for sending text prompts to AI providers.

Supported providers:
- ChatGPT (OpenAI)
- DeepSeek
- Claude (Anthropic)

### Install
```bash
go get github.com/muraduiurie/gpt@latest
```

### Quick start (using config file)
`ai.NewAIAgent` loads settings from `config.yaml` in your app's working directory.

Required keys per provider:
- ChatGPT:
  - `openai_api_token`
  - `openai_text_input_endpoint` (default: `https://api.openai.com/v1/responses` if empty)
- DeepSeek:
  - `deepseek_api_token`
  - `deepseek_text_input_endpoint` (default: `https://api.deepseek.com/chat/completions` if empty)
- Claude:
  - `claude_api_token`
  - `claude_text_input_endpoint` (Anthropic Messages API; default used if empty)

Example `config.yaml`:
```yaml
# ChatGPT
openai_api_token: "YOUR_OPENAI_API_TOKEN"
openai_text_input_endpoint: "https://api.openai.com/v1/responses"

# DeepSeek
deepseek_api_token: "YOUR_DEEPSEEK_API_TOKEN"
deepseek_text_input_endpoint: "https://api.deepseek.com/chat/completions"

# Claude
claude_api_token: "YOUR_CLAUDE_API_TOKEN"
claude_text_input_endpoint: "https://api.anthropic.com/v1/messages"
```

Usage:
```go
package yourpkg

import (
    "github.com/muraduiurie/gpt/pkg/ai"
    cgtypes "github.com/muraduiurie/gpt/pkg/ai/types/chatgpt"
    "github.com/muraduiurie/gpt/pkg/ai/types/union"
)

func ExampleChatGPT() error {
    client, err := ai.NewAIAgent(ai.ChatGPTAgent)
    if err != nil {
        return err
    }

    resp, err := client.AskAI(&union.Request{
        TextRequest: &cgtypes.TextInputRequest{
            Model: string(ai.AiModelGpt4_1),
            Input: "Hello from Go",
        },
    })
    if err != nil {
        return err
    }

    // Type-assert to ChatGPT response type if you need structured fields
    if tr, ok := resp.TextResponse.(*cgtypes.TextInputResponse); ok {
        _ = tr // use tr.Output, tr.Usage, etc.
    }
    return nil
}
```

DeepSeek example:
```go
package yourpkg

import (
    "github.com/muraduiurie/gpt/pkg/ai"
    dstypes "github.com/muraduiurie/gpt/pkg/ai/types/deepseek"
    "github.com/muraduiurie/gpt/pkg/ai/types/union"
)

func ExampleDeepSeek() error {
    client, err := ai.NewAIAgent(ai.DeepSeekAgent)
    if err != nil {
        return err
    }

    resp, err := client.AskAI(&union.Request{
        TextRequest: &dstypes.TextInputRequest{
            Model: string(ai.AiModelDeepSeekChat),
            Messages: []dstypes.TextInputRequestMessage{
                {Role: "user", Content: "Hello from Go"},
            },
        },
    })
    if err != nil {
        return err
    }

    // Type-assert to DeepSeek response type if you need structured fields
    if tr, ok := resp.TextResponse.(*dstypes.TextInputResponse); ok {
        _ = tr // use tr.Choices, tr.Usage, etc.
    }
    return nil
}
```

Claude example:
```go
package yourpkg

import (
    "github.com/muraduiurie/gpt/pkg/ai"
    cltypes "github.com/muraduiurie/gpt/pkg/ai/types/claude"
    "github.com/muraduiurie/gpt/pkg/ai/types/union"
)

func ExampleClaude() error {
    client, err := ai.NewAIAgent(ai.ClaudeAgent)
    if err != nil {
        return err
    }

    resp, err := client.AskAI(&union.Request{
        TextRequest: &cltypes.TextInputRequest{
            Model: cltypes.ClaudeAIModelSonnet4_20250514,
            Messages: []cltypes.TextInputRequestMessage{
                {Role: cltypes.ClaudeAIRoleUser, Content: "Hello from Go"},
            },
        },
    })
    if err != nil {
        return err
    }

    if tr, ok := resp.TextResponse.(*cltypes.TextInputResponse); ok {
        _ = tr // use tr.Content, tr.Usage, etc.
    }
    return nil
}
```

### Manual client (no config file)
If you prefer not to use `config.yaml`, instantiate the client directly:

```go
package yourpkg

import (
    "github.com/muraduiurie/gpt/pkg/ai/providers/chatgpt"
    "github.com/muraduiurie/gpt/pkg/ai/providers/claude"
    "github.com/muraduiurie/gpt/pkg/ai/providers/deepseek"
    cgtypes "github.com/muraduiurie/gpt/pkg/ai/types/chatgpt"
    cltypes "github.com/muraduiurie/gpt/pkg/ai/types/claude"
    dstypes "github.com/muraduiurie/gpt/pkg/ai/types/deepseek"
    "github.com/muraduiurie/gpt/pkg/ai/types/union"
)

func ExampleManualChatGPT() error {
    c := &chatgpt.Client{
        ApiToken:          "YOUR_OPENAI_API_TOKEN",
        TextInputEndpoint: "https://api.openai.com/v1/responses", // optional; defaults if empty
    }
    resp, err := c.AskAI(&union.Request{
        TextRequest: &cgtypes.TextInputRequest{Model: "gpt-4.1", Input: "Ping"},
    })
    if err != nil { return err }
    _ = resp
    return nil
}

func ExampleManualClaude() error {
    c := &claude.Client{
        ApiToken:          "YOUR_CLAUDE_API_TOKEN",
        TextInputEndpoint: "https://api.anthropic.com/v1/messages", // optional; defaults if empty
    }
    resp, err := c.AskAI(&union.Request{
        TextRequest: &cltypes.TextInputRequest{
            Model: cltypes.ClaudeAIModelSonnet4_20250514,
            Messages: []cltypes.TextInputRequestMessage{{Role: cltypes.ClaudeAIRoleUser, Content: "Ping"}},
        },
    })
    if err != nil { return err }
    _ = resp
    return nil
}
```
func ExampleManualDeepSeek() error {
    c := &deepseek.Client{
        ApiToken:          "YOUR_DEEPSEEK_API_TOKEN",
        TextInputEndpoint: "https://api.deepseek.com/chat/completions", // optional; defaults if empty
    }
    resp, err := c.AskAI(&union.Request{
        TextRequest: &dstypes.TextInputRequest{
            Model: "deepseek-chat",
            Messages: []dstypes.TextInputRequestMessage{{Role: "user", Content: "Ping"}},
        },
    })
    if err != nil { return err }
    _ = resp
    return nil
}
```

### Models
Model constants are defined in `github.com/muraduiurie/gpt/pkg/ai`:
- ChatGPT: `AiModelGpt4_1`, `AiModelGpt4o`, `AiModelGpt3_5_turbo`, etc.
- DeepSeek: `AiModelDeepSeekChat`, `AiModelDeepSeekReasoner`.
- Claude (types provide model strings): see `github.com/muraduiurie/gpt/pkg/ai/types/claude` (e.g. `ClaudeAIModelSonnet4_20250514`).

### Types
- Wrapper request/response: `github.com/muraduiurie/gpt/pkg/ai/types/union` (`Request`, `Response`)
- ChatGPT request/response types: `github.com/muraduiurie/gpt/pkg/ai/types/chatgpt`
- DeepSeek request/response types: `github.com/muraduiurie/gpt/pkg/ai/types/deepseek`
- Claude request/response types: `github.com/muraduiurie/gpt/pkg/ai/types/claude`

### Error handling
`AskAI` returns errors for nil inputs, missing required fields, JSON/HTTP failures, and non-2xx responses.



