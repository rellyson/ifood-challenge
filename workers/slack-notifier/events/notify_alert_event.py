import json

class NotifyAlertEvent:
    def __init__(self, sqs_client) -> None:
        self.sqs_client = sqs_client

    def handle(self, queue_url: str):

        print("Started listenning to messages at: {}".format(queue_url))

        messages = self.sqs_client.receive_message(
            QueueUrl=queue_url,
            VisibilityTimeout=0,
            WaitTimeSeconds=0
        )

        for message in messages.get("Messages", []):
            receipt_handle = message['ReceiptHandle']
            message_body = json.loads(message["Body"])

            print("Received Message: {}".format(message_body))

            self.sqs_client.delete_message(
                QueueUrl=queue_url,
                ReceiptHandle=receipt_handle
            )
