package common

import (
	"encoding/json"

	"github.com/invopop/jsonschema"
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
func (Categories) JSONSchemaExtend(schema *jsonschema.Schema) {
	schema.Type = "string"
	schema.Enum = []any{Housing.String(), Groceries.String(), Dining.String(), Transportation.String(), Bills.String(), Utilities.String(), Subscriptions.String(), Entertainment.String(), Shopping.String(), Insurance.String(), Donations.String(), Education.String(), Other.String()}
}

func (cat *Categories) MarshalJSON() ([]byte, error) {

	return json.Marshal(cat.String())

}

func (cat *Categories) UnmarshalJSON(data []byte) error {
	var s string
	var categories = map[string]Categories{
		"housing":        Housing,
		"groceries":      Groceries,
		"dining":         Dining,
		"transportation": Transportation,
		"bills":          Bills,
		"utilities":      Utilities,
		"subscriptions":  Subscriptions,
		"entertainment":  Entertainment,
		"shopping":       Shopping,
		"insurance":      Insurance,
		"donations":      Donations,
		"education":      Education,
		"other":          Other,
	}
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if val, ok := categories[s]; ok {
		*cat = val
		return nil
	}
	return nil

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
