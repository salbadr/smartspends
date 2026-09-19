package ai

import (
	"context"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/responses"
	"github.com/openai/openai-go/shared"
)

var _ AiModel = (*OpenAi)(nil)

type OpenAi struct {
	client *openai.Client
}

func NewOpenAi(apiKey string) *OpenAi {
	c := openai.NewClient(option.WithAPIKey(apiKey))

	return &OpenAi{
		client: &c,
	}
}

func (oa *OpenAi) GetResponse(ctx context.Context, prompt string) string {
	response, err := oa.client.Responses.New(ctx, responses.ResponseNewParams{
		Model: shared.ChatModelGPT4_1Nano,
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String("Tell me a 3 line story about unicorns"),
		},
	})

	if err != nil {
		panic(err.Error())
	}
	return response.OutputText()
}
