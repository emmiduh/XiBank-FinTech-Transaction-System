package main

import (
	"log"
	"net/http"
	"os"
	"ledger-service/db"
	"ledger-service/handlers"
	"ledger-service/middleware"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}
	
	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	// 1. Initialize Database Connection
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	
	// Verify connection is alive immediately
	if err := database.Ping(); err != nil {
		log.Fatalf("Database is unreachable: %v", err)
	}
	defer database.Close()
	log.Println("Successfully connected to Postgres")

	// 2. Initialize Handler with DB dependency
	ledgerHandler := &handlers.LedgerHandler{
		DB: database,
	}

	// 3. Register Routes using the handler instance
	http.Handle("/transaction", middleware.JWTAuth(http.HandlerFunc(ledgerHandler.HandleTransaction)))
	
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	
	log.Println("Ledger service running on port 4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
