package controllers

import (
	"encoding/json"
	"net/http"
)

type HealthCheckController interface {
	Status(w http.ResponseWriter, r *http.Request)
}

type controller struct{}

func NewHealthCheckController() HealthCheckController {
	return &controller{}
}

type HealthcheckResponse struct {
	Message string `json:"message"`
}

func (*controller) Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := HealthcheckResponse{}
	data.Message = "OK"
	json.NewEncoder(w).Encode(data)
}
