package service_test

import (
	"errors"
	"testing"

	awsSQS "github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/rellyson/ifood-challenge/server/service"
	"github.com/rellyson/ifood-challenge/server/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSendingMessage(t *testing.T) {
	sqsClientMock := new(mocks.SQS)
	testService := service.NewMessagingService(sqsClientMock)

	sqsClientMock.On("SQSSendMessage").Return(awsSQS.SendMessageOutput{}, nil)

	res, _ := testService.SendMessageToQueue(service.SendMessagePayload{
		QueueUrl: "test",
		Message:  "test",
	})

	assert.Equal(t, service.SendMessageResponse{
		Status:    "SENT",
		MessageId: "",
		MD5OfBody: "",
	}, res)
}

func TestSendMessagePayloadValidation(t *testing.T) {
	sqsClientMock := new(mocks.SQS)
	testService := service.NewMessagingService(sqsClientMock)

	_, err := testService.SendMessageToQueue(service.SendMessagePayload{
		Message: "test",
	})

	// error assert
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "payload is invalid")
}

func TestSendMessageSQSError(t *testing.T) {
	sqsClientMock := new(mocks.SQS)
	testService := service.NewMessagingService(sqsClientMock)

	errMessage := "sqs test error"
	sqsClientMock.On("SQSSendMessage").Return(awsSQS.SendMessageOutput{}, errors.New(errMessage))

	_, err := testService.SendMessageToQueue(service.SendMessagePayload{
		QueueUrl: "test",
		Message:  "test",
	})

	// error assert
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), errMessage)
}

func TestDeleteMessage(t *testing.T) {
	sqsClientMock := new(mocks.SQS)
	testService := service.NewMessagingService(sqsClientMock)

	sqsClientMock.On("SQSDeleteMessage").Return(awsSQS.DeleteMessageOutput{}, nil)

	err := testService.DeleteMessageFromQueue(service.DeleteMessagePayload{
		QueuUrl:       "test",
		ReceiveHandle: "test",
	})

	assert.Nil(t, err)
}

func TestDeleteMessagePayloadValidation(t *testing.T) {
	sqsClientMock := new(mocks.SQS)
	testService := service.NewMessagingService(sqsClientMock)

	sqsClientMock.On("SQSDeleteMessage").Return(awsSQS.DeleteMessageOutput{}, nil)

	err := testService.DeleteMessageFromQueue(service.DeleteMessagePayload{
		QueuUrl: "test",
	})

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "payload is invalid")
}

func TestDeleteMessageSQSError(t *testing.T) {
	sqsClientMock := new(mocks.SQS)
	testService := service.NewMessagingService(sqsClientMock)

	errMessage := "new test error"
	sqsClientMock.On("SQSDeleteMessage").Return(awsSQS.DeleteMessageOutput{}, errors.New(errMessage))

	err := testService.DeleteMessageFromQueue(service.DeleteMessagePayload{
		QueuUrl:       "test",
		ReceiveHandle: "test",
	})

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), errMessage)
}
