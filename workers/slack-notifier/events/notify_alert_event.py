import json
import asyncio
from services.slack_service import SlackServiceInterface, PostMessagePayload
from aws.sqs_client import SQSClient


class NotifyAlertEvent:
    def __init__(self, sqs_client: SQSClient, slack_service: SlackServiceInterface) -> None:
        self.sqs_client = sqs_client
        self.slack_service = slack_service

    def delivery_incoming_messages(self, queue_url: str, wait_time_secs: int):
        messages = self.sqs_client.receive_message(
            QueueUrl=queue_url,
            MaxNumberOfMessages=1,
            VisibilityTimeout=10,
            WaitTimeSeconds=wait_time_secs
        )

        for message in messages.get("Messages", []):
            receipt_handle = message['ReceiptHandle']
            message_id = message['MessageId']
            message_body = json.loads(message["Body"])

            print("Received Message: {}".format(message_body))

            try:
                slack_message = PostMessagePayload(
                    channel=message_body['channel'], text=message_body['message'],
                    attachments=message_body['attachments'] or None)

                asyncio.run(self.slack_service.post_message(slack_message))

                print("Message {} delivered to channel #{}".format(
                    message_id, message_body['channel']))

                self.sqs_client.delete_message(
                    QueueUrl=queue_url,
                    ReceiptHandle=receipt_handle
                )
            except:
                print("Error handling event message with id {}. Will try again in {} seconds.".format(
                    message_id, wait_time_secs))

    def handle(self, received_signal: bool, queue_url: str, wait_time_secs: int):
        print("Started listening to messages at: {}".format(queue_url))

        while not received_signal:
            self.delivery_incoming_messages(queue_url, wait_time_secs)
