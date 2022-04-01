package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/rellyson/ifood-challenge/server/utils"
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
	b, _ := io.ReadAll(r.Body)

	payload := NotifyAlertPayload{}
	json.Unmarshal(b, &payload)

	err := payload.validate()

	if err != nil {
		utils.ServerError(w, http.StatusBadRequest, "Payload is invalid: "+err.Error())
	}

	w.WriteHeader(http.StatusCreated)
}
