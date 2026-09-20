package expenses

import (
	"encoding/json"
	"net/http"

	"github.com/salbadr/smartspends/internal/common"
)

func ExpensesHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		var er common.ExpensesRequest

		if err := json.NewDecoder(req.Body).Decode(&er); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		defer req.Body.Close()

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		payload := GetExpenses(er.Expenses)
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return http.HandlerFunc(fn)
}
