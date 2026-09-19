package ai

import (
	"context"

	"github.com/anthropics/anthropic-sdk-go"
)
var _ AiModel = (*Claude)(nil)

type Claude struct {
	client *anthropic.Client
}

func (cl *Claude) createClient() anthropic.Client{
	client := anthropic.NewClient()

	return client
}

func (cl *Claude) GetResponse(ctx context.Context, prompt string) string{
	cl.createClient()
	return "claude response"
}
