package expenses

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/salbadr/smartspends/internal/ai"
)

func GetExpenses() ExpenseResponse {

	var resp ExpenseResponse
	resp.Title = "Expense Report"
	resp.Date = time.Now().Format("2006-01-02")
	resp.Transactions = []transaction{
		{
			Date:        "2026-06-01",
			Description: "Toronto Residential Rent",
			Category:    Housing,
			Amount:      2484.45,
		},
		{
			Date:        "2026-06-02",
			Description: "Shoppers Drug Mart",
			Category:    Groceries,
			Amount:      36.45,
		},
		{
			Date:        "2026-06-12",
			Description: "Food Basics",
			Category:    Groceries,
			Amount:      124.53,
		},
		{
			Date:        "2026-06-09",
			Description: "Food Basics",
			Category:    Groceries,
			Amount:      189.18,
		},
		{
			Date:        "2026-06-06",
			Description: "Food Basics",
			Category:    Groceries,
			Amount:      80.90,
		},
	}

	resp.Summary = []summary{
		{Category: Housing, Amount: 2484.45},
		{Category: Groceries, Amount: 536.45},
	}

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, falling back to system environment variables")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is missing")
	}

	c := ai.NewOpenAi(apiKey)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	res := c.GetResponse(ctx, "string")
	log.Println(res)

	defer cancel() // Ensures resources are freed when the function returns

	return resp

}
