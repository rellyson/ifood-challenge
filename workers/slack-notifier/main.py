from logging import fatal, debug
import os
import sys

from aws.sqs_client import new_client
from events.notify_alert_event import NotifyAlertEvent

if __name__ == "__main__":
    print("Starting worker...")

    sqs = new_client()
    queue_url = os.environ.get('SQS_NOTIFY_ALERT_QUEUE')

    if(queue_url == None):
        print("Missing queue_url for notify alert event. Exiting...")
        sys.exit()

    event = NotifyAlertEvent(sqs)
    event.handle(queue_url)
