package handlers

import (
	"encoding/json"
	"net/http"
)

// ResponseWithJSON is a helper function to send JSON responses
func ResponseWithJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}