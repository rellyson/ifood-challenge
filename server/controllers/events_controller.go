package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
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

type NotifyAlertAttachments struct {
	Text       string   `json:"text"`     //required
	Fallback   string   `json:"fallback"` // required
	Blocks     []any    `json:"blocks"`
	Color      string   `json:"color"`
	Value      string   `json:"value"`
	Short      string   `json:"short"`
	AuthorIcon string   `json:"author_icon"`
	AuthorLink string   `json:"author_link"`
	AuthorName string   `json:"author_name"`
	Fields     []any    `json:"fields"`
	Footer     string   `json:"footer"`
	FooterIcon string   `json:"footer_icon"`
	ImageUrl   string   `json:"image_url"`
	MrkdwnIn   []string `json:"mrkdwn_in"`
	Pretext    string   `json:"pretext"`
	ThumbUrl   string   `json:"thumb_url"`
	Title      string   `json:"title"`
	TitleLink  string   `json:"title_link"`
	Ts         int64    `json:"ts"`
}

type NotifyAlertPayload struct {
	Channel     string                   `json:"channel"`
	Message     any                      `json:"message"`
	Attachments []NotifyAlertAttachments `json:"attachments"`
}

func (p *NotifyAlertPayload) validate() error {
	if p.Channel == "" {
		return errors.New("channel field is missing")
	}

	if p.Message == nil {
		return errors.New("message field is missing")
	}

	for index, element := range p.Attachments {
		if element.Text == "" {
			return errors.New(fmt.Sprintf("attachment on index %v is missing field text", index))
		}

		if element.Fallback == "" {
			return errors.New(fmt.Sprintf("attachment on index %v is missing field fallback", index))
		}
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

	serviceRes, serviceErr := messageService.SendMessageToQueue(service.SendMessagePayload{
		QueueUrl: os.Getenv("SQS_NOTIFY_QUEUE"),
		Message:  payload,
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
