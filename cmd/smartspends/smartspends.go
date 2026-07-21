package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
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

type ExpenseRequest struct {
	Expense string `json:"expense"`
}

type Transaction struct {
	Date        string     `json:"date"`
	Description string     `json:"description"`
	Category    Categories `json:"category"`
	Amount      float32    `json:"amount"`
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

func (cat Categories) MarshalJSON() ([]byte, error) {

	return json.Marshal(cat.String())

}

func main() {
	var expense = `Date,Description,Amount,Balance
2026-06-01,Toronto Residential Rent,-1850.0,3582.1
2026-06-02,Metro Supermarket,-47.25,3534.85
2026-06-02,Amazon.ca,-77.22,3457.63
2026-06-03,TTC Toronto Transit,-3.35,3454.28
2026-06-05,Metro Supermarket,-57.87,3396.41
2026-06-06,Toronto Hydro,-92.44,3303.97
2026-06-07,TTC Toronto Transit,-3.35,3300.62
2026-06-10,Uber Eats,-38.64,3261.98
2026-06-11,Starbucks Coffee,-6.64,3255.34
2026-06-12,Metro Supermarket,-43.05,3212.29
2026-06-12,Toronto Hydro,-93.71,3118.58
2026-06-13,Loblaws Toronto,-134.63,2983.95
2026-06-15,Direct Deposit / Payroll,2850.0,5833.95
2026-06-15,Apple.com/bill,-12.99,5820.96
2026-06-15,Sobeys,-77.69,5743.27
2026-06-16,TTC Toronto Transit,-3.35,5739.92
2026-06-16,Uber Trip,-15.73,5724.19
2026-06-17,TTC Toronto Transit,-3.35,5720.84
2026-06-17,Toronto Hydro,-88.99,5631.85
2026-06-18,Starbucks Coffee,-7.5,5624.35
2026-06-18,Starbucks Coffee,-6.68,5617.67
2026-06-20,Amazon.ca,-16.64,5601.03
2026-06-21,Uber Eats,-48.29,5552.74
2026-06-22,Tim Hortons,-7.74,5545.0
2026-06-24,Sobeys,-60.19,5484.81
2026-06-24,Uber Trip,-29.86,5454.95
2026-06-28,Apple.com/bill,-12.99,5441.96
2026-06-29,Metro Supermarket,-52.4,5389.56
2026-06-30,Direct Deposit / Payroll,2850.0,8239.56`
	message := fmt.Sprintf("{\"expense\": \"%s\"}", strings.ReplaceAll(expense, "\n", ""))

	req := []byte(message)
	var er ExpenseRequest

	if err := json.Unmarshal(req, &er); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println("THE REQUEST IS:")
	fmt.Println(er)

	var resp ExpenseResponse
	resp.Title = "Expense Report"
	resp.Date = time.Now().Format("2006-01-02")
	resp.Transactions = []Transaction{
		{
			Date:        "2026-06-01",
			Description: "Toronto Residential Rent",
			Category:    Housing,
			Amount:      123.45,
		},
		{
			Date:        "2026-06-02",
			Description: "Shoppers Drug Mart",
			Category:    Groceries,
			Amount:      36.45,
		},
	}

	resp.Summary = []Summary{
		{Category: Housing, Amount: 123.45},
		{Category: Groceries, Amount: 36.45},
	}

	res, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error Marshalling JSON:", err)
		return
	}

	fmt.Println("THE RESPONSE IS:")
	fmt.Println(string(res))

}
