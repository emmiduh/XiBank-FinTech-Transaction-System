package handlers
import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// LedgerHandler holds dependencies for the handlers
type LedgerHandler struct {
	DB *sql.DB
}

type Transaction struct {
	UserID string `json:"user_id"`
	Amount float64 `json:"amount"`
}

func (h *LedgerHandler) HandleTransaction(w http.ResponseWriter, r *http.Request) {
	var tx Transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if tx.Amount <= 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	// Retrieve UserID from context (set by your JWT middleware)
	userIDHeader := r.Header.Get("X-User-ID")
	if userIDHeader != "" {
		tx.UserID = userIDHeader
	}

	// --- DATABASE INTERACTION ---
	// Ensure you have a 'transactions' table created in your Postgres DB
	query := `INSERT INTO transactions (user_id, amount) VALUES ($1, $2) RETURNING id`
	var newID int
	err := h.DB.QueryRow(query, tx.UserID, tx.Amount).Scan(&newID)
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":         "transaction approved",
		"transaction_id": newID,
	})
}