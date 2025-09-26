## Go AI Client Library

Lightweight Go client for sending text prompts to AI providers (currently ChatGPT).

### Install
```bash
go get github.com/muraduiurie/gpt@latest
```

### Quick start (using config file)
`ai.NewAIAgent` loads settings from `config.yaml` in your app's working directory.

Required keys:
- `api_token`: API key/token
- `text_input_endpoint`: text responses endpoint (e.g. `https://api.openai.com/v1/responses`)

Example `config.yaml`:
```yaml
api_token: "YOUR_API_TOKEN"
text_input_endpoint: "https://api.openai.com/v1/responses"
```

Usage:
```go
package yourpkg

import (
    "github.com/muraduiurie/gpt/pkg/ai"
    cgtypes "github.com/muraduiurie/gpt/pkg/ai/types/chatgpt"
    "github.com/muraduiurie/gpt/pkg/ai/types/common"
)

func Example() error {
    client, err := ai.NewAIAgent(ai.ChatGPTAgent)
    if err != nil {
        return err
    }

    resp, err := client.AskAI(&common.Request{
        TextRequest: &cgtypes.TextInputRequest{
            Model: string(ai.AiModelGpt4_1),
            Input: "Hello from Go",
        },
    })
    if err != nil {
        return err
    }

    _ = resp // handle response
    return nil
}
```

### Manual client (no config file)
If you prefer not to use `config.yaml`, instantiate the client directly:

```go
package yourpkg

import (
    "github.com/muraduiurie/gpt/pkg/chatgpt"
    "github.com/muraduiurie/gpt/pkg/ai/types/common"
    cgtypes "github.com/muraduiurie/gpt/pkg/ai/types/chatgpt"
)

func ExampleManual() error {
    c := &chatgpt.Client{
        ApiToken:          "YOUR_API_TOKEN",
        TextInputEndpoint: "https://api.openai.com/v1/responses",
    }

    resp, err := c.AskAI(&common.Request{
        TextRequest: &cgtypes.TextInputRequest{
            Model: "gpt-4.1",
            Input: "Ping",
        },
    })
    if err != nil {
        return err
    }
    _ = resp
    return nil
}
```

### Models
Model constants are defined in `github.com/muraduiurie/gpt/pkg/ai` (e.g. `AiModelGpt4_1`, `AiModelGpt4o`, `AiModelGpt3_5_turbo`).

### Types
- Request type: `github.com/muraduiurie/gpt/pkg/ai/types/common.Request`
- ChatGPT request/response types: `github.com/muraduiurie/gpt/pkg/ai/types/chatgpt`

### Error handling
`AskAI` returns errors for nil inputs, empty message/model, JSON/HTTP failures, and non-2xx responses.



