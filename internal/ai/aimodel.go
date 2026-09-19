package ai

import "context"

type AiModel interface{
	GetResponse(ctx context.Context, prompt string) string
}