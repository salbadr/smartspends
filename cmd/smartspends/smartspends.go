package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/salbadr/smartspends/internal/expenses"
)

func main() {
	log.Println("Starting")
	

	mux := http.NewServeMux()
	expenses.RegisterExpenseRoutes(mux)

	mux.Handle("/", http.NotFoundHandler())
	log.Println("Server started on port 8000")

	handler := cors.Default().Handler(mux)

	if err := http.ListenAndServe(":8000", handler); err != nil {
		log.Fatal(err)
	}

}
