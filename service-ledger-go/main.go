package main

import (
	"log"
	"net/http"
	"os"
	"ledger-service/handlers"
	"ledger-service/middleware"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}
	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	http.Handle("/transaction", middleware.JWTAuth(http.HandlerFunc(handlers.HandleTransaction)))
	log.Println("Ledger service running on port 4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}