package expenses

import (
	"net/http"
)


func RegisterExpenseRoutes(router *http.ServeMux) {
	router.Handle("POST /api/v1/expenses", ExpensesHandler())
}
