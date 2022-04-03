package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/rellyson/ifood-challenge/server/aws"
)

type MessagingService interface {
	SendMessageToQueue(p SendMessagePayload) (SendMessageResponse, error)
	DeleteMessageFromQueue(p DeleteMessagePayload) error
}

type service struct{}

var (
	sqsClient aws.SQSClient
)

type SendMessagePayload struct {
	QueueUrl string
	Message  any
}

type SendMessageResponse struct {
	Status    string `json:"status"`
	MessageId string `json:"message_id"`
	MD5OfBody string `json:"md5_of_body"`
}

type DeleteMessagePayload struct {
	QueuUrl       string
	ReceiveHandle string
}

func NewMessagingService(client aws.SQSClient) MessagingService {
	sqsClient = client
	return &service{}
}

func (*service) SendMessageToQueue(p SendMessagePayload) (SendMessageResponse, error) {
	if p.QueueUrl == "" || p.Message == nil {
		err := errors.New("payload is invalid")
		return SendMessageResponse{}, err
	}

	if reflect.TypeOf(p.Message).Kind() != reflect.String {
		s, _ := json.Marshal(p.Message)
		p.Message = string(s)
	}

	r, err := sqsClient.SQSSendMessage(p.QueueUrl, fmt.Sprintf("%v", p.Message))

	if err != nil {
		return SendMessageResponse{}, err
	}

	return SendMessageResponse{
		Status:    "SENT",
		MessageId: *r.MessageId,
		MD5OfBody: *r.MD5OfMessageBody,
	}, nil
}

func (*service) DeleteMessageFromQueue(p DeleteMessagePayload) error {
	if p.QueuUrl == "" || p.ReceiveHandle == "" {
		err := errors.New("payload is invalid")
		return err
	}

	_, err := sqsClient.SQSDeleteMessage(p.QueuUrl, p.ReceiveHandle)

	if err != nil {
		return err
	}

	return nil
}
