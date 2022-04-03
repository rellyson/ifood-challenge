package mocks

import (
	awsSQS "github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/stretchr/testify/mock"
)

type SQS struct {
	mock.Mock
}

func (m *SQS) SQSSendMessage(queueUrl string, message string) (awsSQS.SendMessageOutput, error) {
	args := m.Called()

	return awsSQS.SendMessageOutput{}, args.Error(1)
}

func (m *SQS) SQSDeleteMessage(queueUrl string, rcvHnd string) (awsSQS.DeleteMessageOutput, error) {
	args := m.Called()

	return awsSQS.DeleteMessageOutput{}, args.Error(1)
}
