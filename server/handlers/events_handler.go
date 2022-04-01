package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type NotifyAlertPayload struct {
	Channel string `json:"channel"`
	Message any    `json:"message"`
}

func (p *NotifyAlertPayload) validate() error {
	if p.Channel == "" {
		return errors.New("channel field is missing")
	}

	if p.Message == nil {
		return errors.New("message field is missing")
	}
	return nil
}

func NotifyAlertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	b, _ := io.ReadAll(r.Body)

	if b == nil {
		http.Error(w, "Missing request body", http.StatusBadRequest)
		return
	}

	payload := NotifyAlertPayload{}
	json.Unmarshal(b, &payload)

	err := payload.validate()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
