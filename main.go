package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/muraduiurie/gpt/pkg/ai"
	cgtypes "github.com/muraduiurie/gpt/pkg/ai/types/chatgpt"
	cltypes "github.com/muraduiurie/gpt/pkg/ai/types/claude"
	dstypes "github.com/muraduiurie/gpt/pkg/ai/types/deepseek"
	"github.com/muraduiurie/gpt/pkg/ai/types/union"
)

func main() {
	// chatgpt
	openai, err := ai.NewAIAgent(ai.ChatGPTAgent)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	resp, err := openai.AskAI(&union.Request{
		TextRequest: &cgtypes.TextInputRequest{
			Model: cgtypes.AiModelGpt4_1,
			Input: "Hey, this is a test message",
		},
	})
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println("marshal error:", err)
		return
	}
	fmt.Println(string(b))

	// deepseek
	ds, err := ai.NewAIAgent(ai.DeepSeekAgent)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	resp, err = ds.AskAI(&union.Request{
		TextRequest: &dstypes.TextInputRequest{
			Model: dstypes.DeepSeekAIModelChat,
			Messages: []dstypes.TextInputRequestMessage{
				{
					Role:    dstypes.DeepSeekAIRoleUser,
					Content: "Hey, this is a test message",
				},
			},
		},
	})
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	b, err = json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println("marshal error:", err)
		return
	}
	fmt.Println(string(b))

	// claude
	cl, err := ai.NewAIAgent(ai.ClaudeAgent)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	resp, err = cl.AskAI(&union.Request{
		TextRequest: &cltypes.TextInputRequest{
			Model:     cltypes.ClaudeAIModelSonnet4_20250514,
			MaxTokens: 100,
			Messages: []cltypes.TextInputRequestMessage{
				{
					Role:    cltypes.ClaudeAIRoleUser,
					Content: "Hey, this is a test message",
				},
			},
		},
	})
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	b, err = json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println("marshal error:", err)
		return
	}
	fmt.Println(string(b))
}
