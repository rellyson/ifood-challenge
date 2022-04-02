package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthcheckResponse struct {
	Message string `json:"message"`
}

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := HealthcheckResponse{}
	data.Message = "OK"
	json.NewEncoder(w).Encode(data)
}
