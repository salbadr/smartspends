package expenses

import (
	"encoding/json"
	"net/http"
)

func expensesHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		var er ExpenseRequest

		if err := json.NewDecoder(req.Body).Decode(&er); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		defer req.Body.Close()

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		payload := GetExpenses()
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return http.HandlerFunc(fn);
}

func RegisterExpenseRoutes(router *http.ServeMux) {
	router.Handle("POST /api/v1/expenses", expensesHandler())
}
