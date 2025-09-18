package main

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/ark"
	model2 "github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  os.Getenv("MODEL"),
	})
	if err != nil {
		panic(err)
	}
	input := []*schema.Message{
		schema.SystemMessage("You are a helpful assistant."),
		schema.UserMessage("What is the meaning of life?"),
	}
	reader, err := model.Stream(ctx, input, model2.WithMaxTokens(64))
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	for {
		chunk, err := reader.Recv()
		if err != nil {
			panic(err)
		}
		print(chunk.Content)
	}

}
