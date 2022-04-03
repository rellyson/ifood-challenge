package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/rellyson/ifood-challenge/server/service"
	"github.com/rellyson/ifood-challenge/server/utils"
)

type EventsController interface {
	NotifyAlert(w http.ResponseWriter, r *http.Request)
}

type handler struct{}

var (
	messageService service.MessagingService
)

func NewEventsController(service service.MessagingService) EventsController {
	messageService = service
	return &handler{}
}

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

func (*handler) NotifyAlert(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)

	payload := NotifyAlertPayload{}
	json.Unmarshal(b, &payload)

	validationErr := payload.validate()

	if validationErr != nil {
		utils.ServerError(w, http.StatusBadRequest, "Payload is invalid: "+validationErr.Error())
		return
	}

	alertMessage := struct {
		Channel string `json:"channel"`
		Message any    `json:"message"`
	}{
		payload.Channel,
		payload.Message,
	}

	serviceRes, serviceErr := messageService.SendMessageToQueue(service.SendMessagePayload{
		QueueUrl: os.Getenv("SQS_NOTIFY_QUEUE"),
		Message:  alertMessage,
	})

	if serviceErr != nil {
		log.Printf("Error when sending message: %v", serviceErr)
		utils.ServerError(w, http.StatusInternalServerError, "Something went wrong. try again later")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(service.SendMessageResponse{
		Status:    serviceRes.Status,
		MessageId: serviceRes.MessageId,
		MD5OfBody: serviceRes.MD5OfBody,
	})
}
