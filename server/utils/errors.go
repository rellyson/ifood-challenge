package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func ServerError(r http.ResponseWriter, s int, err string) ErrorResponse {
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(s)

	e := ErrorResponse{}
	e.Status = s
	e.Error = err

	json.NewEncoder(r).Encode(e)
	return e
}
