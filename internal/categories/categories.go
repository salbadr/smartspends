package categories

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