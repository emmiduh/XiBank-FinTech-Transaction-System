package main

import (
	"log"
	"net/http"
	"ledger-service/handlers"
)

func main() {
	http.HandleFunc("/transaction", handlers.HandleTransaction)
	log.Println("Ledger service running on port 4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}