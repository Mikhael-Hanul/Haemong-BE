package service

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type AiService struct {
	token string
}

func NewAiService(token string) *AiService {
	return &AiService{
		token: token,
	}
}

func (r *AiService) Haemong(content string) (string, error) {
	c := openai.NewClient(r.token)

	resp, err := c.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model:       openai.GPT3Dot5TurboInstruct,
			Prompt:      "다음 꿈에 대한 해몽을 해주세요: " + content,
			MaxTokens:   1024,
			Temperature: 0.7,
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Text, nil
}
