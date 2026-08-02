package main

import (
	"log"
	"github.com/salbadr/smartspends/internal/expenses"
	"net/http"
)


func main() {
	
	mux := http.NewServeMux()
	expenses.RegisterExpenseRoutes(mux)
	mux.Handle("/", http.NotFoundHandler())
	log.Println("Server started on port 8000")

	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}

}
