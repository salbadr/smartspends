package ai

import (
	"context"
	"encoding/json"
	"log"

	"github.com/invopop/jsonschema"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/responses"
	"github.com/openai/openai-go/shared"
	"github.com/salbadr/smartspends/internal/common"
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

func (oa *OpenAi) getSchema() map[string]any {
	reflector := &jsonschema.Reflector{Anonymous: true}
	schema := reflector.Reflect(&common.ExpenseResponse{})
	// encode the schema to a human readable format
	schemaJSON, err := json.Marshal(schema)
	if err != nil {
		log.Fatalf("Failed to marshal schema: %v", err)
	}

	//decode the schema as map[string]any.
	var dat map[string]any
	if err := json.Unmarshal(schemaJSON, &dat); err != nil {
		panic(err)
	}

	return dat

}

func (oa *OpenAi) GetResponse(ctx context.Context, prompt string) string {
	response, err := oa.client.Responses.New(ctx, responses.ResponseNewParams{
		Model: shared.ChatModelGPT4_1Nano,
		Instructions: openai.String(
			`You are a professional financial accountant. You will be provided
			with an input containing the credit card expenses. 
			
			Your task is to:
			1. Categorize the expenses. You will be provided with the list of categories 
			2. Reason step by step before categorizing. 
				An expense to "KANDAHAR KABAB" should be in "Dining" and not "Grocery" because it is a restaurant.
			3. The Title of the report should have the title "Expense Report" in it.	
			4. Finally, create a summary that has the total expenditure per category 
		`),
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String(prompt),
		},
		Text: responses.ResponseTextConfigParam{
			Format: responses.ResponseFormatTextConfigUnionParam{
				OfJSONSchema: &responses.ResponseFormatTextJSONSchemaConfigParam{
					Name:   "expense_reasoning",
					Schema: oa.getSchema(),
					Strict: openai.Bool(true),
				},
			},
		},
	})

	if err != nil {
		panic(err.Error())
	}
	return response.OutputText()
}
