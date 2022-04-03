package aws

import (
	"context"

	awsSQS "github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSClient interface {
	SQSSendMessage(queueUrl string, message string) (awsSQS.SendMessageOutput, error)
	SQSDeleteMessage(queueUrl string, rcvHnd string) (awsSQS.DeleteMessageOutput, error)
}

type client struct{}

var (
	sqs *awsSQS.Client
)

func NewSQSClient() SQSClient {
	sqs = awsSQS.NewFromConfig(AwsConfig())
	return &client{}
}

func (*client) SQSSendMessage(queueUrl string, message string) (awsSQS.SendMessageOutput, error) {
	r, err := sqs.SendMessage(context.TODO(), &awsSQS.SendMessageInput{
		QueueUrl:    &queueUrl,
		MessageBody: &message,
	})

	if err != nil {
		return awsSQS.SendMessageOutput{}, err
	}

	return awsSQS.SendMessageOutput{
		MessageId:        r.MessageId,
		MD5OfMessageBody: r.MD5OfMessageBody,
	}, nil
}

func (*client) SQSDeleteMessage(queueUrl string, rcvHnd string) (awsSQS.DeleteMessageOutput, error) {
	r, err := sqs.DeleteMessage(context.TODO(), &awsSQS.DeleteMessageInput{
		QueueUrl:      &queueUrl,
		ReceiptHandle: &rcvHnd,
	})

	if err != nil {
		return awsSQS.DeleteMessageOutput{}, err
	}

	return awsSQS.DeleteMessageOutput{
		ResultMetadata: r.ResultMetadata,
	}, nil
}
