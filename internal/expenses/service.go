package expenses

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/salbadr/smartspends/internal/ai"
	"github.com/salbadr/smartspends/internal/common"
)

func GetExpenses(payload string) *common.ExpenseResponse {


	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, falling back to system environment variables")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is missing")
	}

	c := ai.NewOpenAi(apiKey)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	res := c.GetResponse(ctx, payload)

	resp := common.ExpenseResponse{}
	json.Unmarshal([]byte(res), &resp)

	defer cancel() // Ensures resources are freed when the function returns

	return &resp

}
