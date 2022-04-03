<p align="center">
<img src="./assets/Ifood-logo.png" width="200">
</p>

# iFood Assessment Challenge

## Description

As we receive security alerts from different systems and, and it's not viable checking incidents manually, **we need an automated way to receive them** in *Slack* and **ensure that no alert will be lost due to the unavailability of some component.**
Therefore, your objective in this challenge will be to **develop a system that receives alerts and notifies them in different channels in *Slack***, guaranteeing the delivery of messages through an *AWS SQS* queue.

Your system **must have the following components**:
- A server that receives JSON events through its API and sends them to an AWS SQS queue. These events must at least contain the message and the *Slack* channel where that message will be posted.
- An AWS SQS queue.
- A worker that consumes these messages from the queue and post them to the *Slack* channel.

## Project requirements

- The server must be implemented in Go and the worker in Python.
- The server, worker and AWS environment must be dockerized.
- Documentation and testing.

## Running the project

To run the project you have these alternatives:

``` bash
# Runs using base image
$ docker-compose up

# Runs in a development environment, enabling auto reload
$ docker-compose -f .docker/dev-compose.yaml up

```

## Sending an alert notification

To send an alert notification to a Slack channel through this system, send a request to `/v1/events/notify_alert` endpoint following the payload specified below.

``` JSON
{
    "channel": "my-notification-channel",
    "messsage": "Hi from notification server"
}
```

Example:

``` bash

# Using curl
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"channel":"my-notification-channel","message":"This is a message"}' \
  http://localhost:3000/v1/events/notify_alert

```

## How alert notification works

This diagram ilustrates the notify alert sequence flow.

<img src="./assets/notify-alert-sequence.png"/>

## Useful links

- [PlantUML](https://plantuml.com/)
- [AWS SDK for Go](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/)
- [AWS SDK for Python](https://boto3.amazonaws.com/v1/documentation/api/latest/index.html)
- [Boto3 Guide](https://boto3.amazonaws.com/v1/documentation/api/1.21.30/guide/sqs-examples.html)