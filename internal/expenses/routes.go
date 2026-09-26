package expenses

import (
	"errors"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/salbadr/smartspends/internal/ai"
)

func RegisterExpenseRoutes(router *http.ServeMux) error {
	if err := godotenv.Load(); err != nil {
		return errors.New("Warning: No .env file found")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return errors.New("Failed to obtain API key")

	}

	aiClient, err := ai.NewOpenAi(apiKey)
	if err != nil {
		return errors.New("Could not instantiate AI Client")

	}

	router.Handle("POST /api/v1/expenses", ExpensesHandler(aiClient))
	return nil
}
