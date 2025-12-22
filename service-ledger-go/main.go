package main

import (
	"log"
	"net/http"
	"ledger-service/handlers"
)

func main() {
	http.HnaldeFunc("/transaction", handlers.HandlerTransaction)
	log.Println("Ledeger service running on port 4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}