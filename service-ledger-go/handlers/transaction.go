package handlers
import (
	"encoding/json"
	"net/http"
)

type Transaction struct {
	UserID string `json":user_id"`
	Amount float64 `json:"amount"`
}

funct HandleTransaction(w http.ResponseWriter, r *http.Request) {
	var tx Transaction
	json.NewDecoder(r.Body).Decode(&tx)

	if tx.Amount <= 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "transaction approved",
	})
}