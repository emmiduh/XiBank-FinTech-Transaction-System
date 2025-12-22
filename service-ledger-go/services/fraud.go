package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func CheckFraud(amount float64) (bool, error) {
	body, _ := json.Marshal(map[string]float64 {
		"amount": amount,
	})

	resp, err := http.Post(
		"http://fraud:5000/check",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return true, err
	}
	defer resp.Body.Close()

	var result map[string]bool
	json.NewDecoder(resp.Body).Decode(&result)

	return result["fraud"], nil
}