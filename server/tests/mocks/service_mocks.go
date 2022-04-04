package mocks

import (
	"github.com/rellyson/ifood-challenge/server/service"
	"github.com/stretchr/testify/mock"
)

type MessageService struct {
	mock.Mock
}

func (m *MessageService) SendMessageToQueue(p service.SendMessagePayload) (service.SendMessageResponse, error) {
	args := m.Called()

	return service.SendMessageResponse{}, args.Error(1)
}

func (m *MessageService) DeleteMessageFromQueue(p service.DeleteMessagePayload) error {
	args := m.Called()

	return args.Error(1)
}
