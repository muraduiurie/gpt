package main

import (
	"encoding/json"
	"fmt"

	"github.com/muraduiurie/gpt/pkg/ai"
	cgtypes "github.com/muraduiurie/gpt/pkg/ai/types/chatgpt"
	"github.com/muraduiurie/gpt/pkg/ai/types/common"
)

func main() {
	c, err := ai.NewAIAgent(ai.ChatGPTAgent)
	if err != nil {
		panic(err)
	}

	resp, err := c.AskAI(&common.Request{
		TextRequest: &cgtypes.TextInputRequest{
			Model: string(ai.AiModelGpt4_1),
			Input: "Hey, this is a test message",
		},
	})
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println("marshal error:", err)
		return
	}
	fmt.Println(string(b))
}
