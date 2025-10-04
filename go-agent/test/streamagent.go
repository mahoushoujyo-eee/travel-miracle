package main

import (
	"log"

	"github.com/cloudwego/eino/adk"
)

func main3(){
	input := adk.AgentInput{
		Messages: []adk.Message{
			{
				Name: "user",
				Role: "user",
				Content: "请为我规划下从南京到潍坊的出行路线，谢谢！",
			},
		},
		EnableStreaming: true,
	}

	log.Printf("input: %v", input)
}