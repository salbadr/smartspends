package common

import (
	"encoding/json"
)

type Categories int

const (
	Housing Categories = iota
	Groceries
	Dining
	Transportation
	Bills
	Utilities
	Subscriptions
	Entertainment
	Shopping
	Insurance
	Donations
	Education
	Other
)

func (c Categories) String() string {
	var categories = map[Categories]string{
		Housing:        "housing",
		Groceries:      "groceries",
		Dining:         "dining",
		Transportation: "transportation",
		Bills:          "bills",
		Utilities:      "utilities",
		Subscriptions:  "subscriptions",
		Entertainment:  "entertainment",
		Shopping:       "shopping",
		Insurance:      "insurance",
		Donations:      "donations",
		Education:      "education",
		Other:          "other",
	}

	if str, ok := categories[c]; ok {
		return str
	}
	return "unknown"
}

func (cat *Categories) MarshalJSON() ([]byte, error) {

	return json.Marshal(cat.String())

}

type ExpensesRequest struct {
	Expenses string `json:"expenses"`
}

type Transaction struct {
	Date        string     `json:"date"`
	Description string     `json:"description"`
	Category    Categories `json:"category"`
	Amount      float32    `json:"amount"`
	Confidence  float32    `json:"confidence"`
}

type Summary struct {
	Category Categories `json:"category"`
	Amount   float32    `json:"amount"`
}

type ExpenseResponse struct {
	Title        string        `json:"title"`
	Date         string        `json:"date"`
	Transactions []Transaction `json:"transactions"`
	Summary      []Summary     `json:"summary"`
}
