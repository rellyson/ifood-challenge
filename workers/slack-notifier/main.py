from logging import fatal, debug
import os
from re import S
import sys

from aws.sqs_client import new_client
from events.notify_alert_event import NotifyAlertEvent
from services.slack_service import SlackService

if __name__ == "__main__":
    print("Starting worker...")

    sqs = new_client()
    queue_url = os.environ.get('SQS_NOTIFY_ALERT_QUEUE')

    if(queue_url == None):
        print("Missing queue_url for notify alert event. Exiting...")
        sys.exit()

    slack = SlackService()

    event = NotifyAlertEvent(sqs, slack)
    event.handle(queue_url)
