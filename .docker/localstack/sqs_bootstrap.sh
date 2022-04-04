#!/usr/bin/env bash

export LOCALSTACK_ENDPOINT=http://localhost:4566
export AWS_DEFAULT_REGION=us-east-1
export AWS_ACCESS_KEY_ID=foobar
export AWS_SECRET_ACCESS_KEY=foobar

# Create queues

# Dead letter queue
aws --endpoint-url=$LOCALSTACK_ENDPOINT sqs create-queue --queue-name notification_dlx --region $AWS_DEFAULT_REGION # DLX queue
DLQ_SQS_ARN=$(aws --endpoint-url=$LOCALSTACK_ENDPOINT sqs get-queue-attributes\
                  --attribute-name QueueArn --queue-url=$LOCALSTACK_ENDPOINT/000000000000/notification_dlx\
                  |  sed 's/"QueueArn"/\n"QueueArn"/g' | grep '"QueueArn"' | awk -F '"QueueArn":' '{print $2}' | tr -d '"' | xargs)


# event queue 
aws --endpoint-url=$LOCALSTACK_ENDPOINT sqs create-queue --queue-name notify_alert_event --region $AWS_DEFAULT_REGION \
    --attributes '{
                    "RedrivePolicy": "{\"deadLetterTargetArn\":\"'"$DLQ_SQS_ARN"'\",\"maxReceiveCount\":\"10\"}"
                    }'