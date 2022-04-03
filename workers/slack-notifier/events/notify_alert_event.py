import json
import asyncio
from services.slack_service import SlackServiceInterface, MessagePayload


class NotifyAlertEvent:
    def __init__(self, sqs_client, slack_service: SlackServiceInterface) -> None:
        self.sqs_client = sqs_client
        self.slack_service = slack_service

    def handle(self, queue_url: str):

        print("Started listenning to messages at: {}".format(queue_url))

        while True:
            messages = self.sqs_client.receive_message(
                QueueUrl=queue_url,
                MaxNumberOfMessages=1,
                VisibilityTimeout=10,
                WaitTimeSeconds=10
            )

            for message in messages.get("Messages", []):
                receipt_handle = message['ReceiptHandle']
                message_body = json.loads(message["Body"])

                print("Received Message: {}".format(message_body))

                try:
                    slack_message = MessagePayload(
                        channel=message_body['channel'], text=message_body['message'])

                    asyncio.run(self.slack_service.post_message(slack_message))

                    print("Message {} delivered to channel: {}".format(
                        message['MessageId'], message_body['channel']))

                    self.sqs_client.delete_message(
                        QueueUrl=queue_url,
                        ReceiptHandle=receipt_handle
                    )
                except:
                    print("Error handling event message: Will try again in 10 seconds.")
