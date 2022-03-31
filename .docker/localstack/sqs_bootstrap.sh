#!/usr/bin/env bash

export LOCALSTACK_ENDPOINT=http://localhost:4566
export AWS_DEFAULT_REGION=us-east-1
export AWS_ACCESS_KEY_ID=foobar
export AWS_SECRET_ACCESS_KEY=foobar

# Create queues
aws --endpoint-url=$LOCALSTACK_ENDPOINT sqs create-queue --queue-name notify_alert_event --region $AWS_DEFAULT_REGION