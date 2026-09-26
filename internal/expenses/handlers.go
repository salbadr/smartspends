package expenses

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/salbadr/smartspends/internal/ai"
	"github.com/salbadr/smartspends/internal/common"
)

func ExpensesHandler(aiClient ai.AiModel) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		var er common.ExpensesRequest

		if err := json.NewDecoder(req.Body).Decode(&er); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		defer req.Body.Close()

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		payload, err := GetExpenses(er.Expenses, aiClient)
		if err != nil {
			log.Printf("Unable to generate response: %v", err)
		}
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return http.HandlerFunc(fn)
}
