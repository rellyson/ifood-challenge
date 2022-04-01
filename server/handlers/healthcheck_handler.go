package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthcheckResponse struct {
	Message string `json:"message"`
}

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	data := HealthcheckResponse{}
	data.Message = "OK"
	json.NewEncoder(w).Encode(data)
	w.WriteHeader(http.StatusOK)
}
