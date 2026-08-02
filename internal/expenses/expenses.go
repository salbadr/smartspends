package expenses
import (
	"time"
	"github.com/salbadr/smartspends/internal/categories"
)


type ExpenseRequest struct {
	Expense string `json:"expense"`
}

type transaction struct {
	Date        string     `json:"date"`
	Description string     `json:"description"`
	Category    categories.Categories `json:"category"`
	Amount      float32    `json:"amount"`
}

type summary struct {
	Category categories.Categories `json:"category"`
	Amount   float32    `json:"amount"`
}

type ExpenseResponse struct {
	Title        string        `json:"title"`
	Date         string        `json:"date"`
	Transactions []transaction `json:"transactions"`
	Summary      []summary     `json:"summary"`
}


func GetExpenses() ExpenseResponse {
	var resp ExpenseResponse
	resp.Title = "Expense Report"
	resp.Date = time.Now().Format("2006-01-02")
	resp.Transactions = []transaction{
		{
			Date:        "2026-06-01",
			Description: "Toronto Residential Rent",
			Category:    categories.Housing,
			Amount:      123.45,
		},
		{
			Date:        "2026-06-02",
			Description: "Shoppers Drug Mart",
			Category:    categories.Groceries,
			Amount:      36.45,
		},
	}

	resp.Summary = []summary{
		{Category: categories.Housing, Amount: 123.45},
		{Category: categories.Groceries, Amount: 36.45},
	}

	return resp

}
