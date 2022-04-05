import os
from slack_sdk import WebClient
from utils.signal_handler import SignalHandler
import sys

from aws.sqs_client import SQSClient
from events.notify_alert_event import NotifyAlertEvent
from services.slack_service import SlackService

if __name__ == "__main__":
    print("Starting worker...")
    queue_url = os.environ.get('SQS_NOTIFY_ALERT_QUEUE')
    oauth_token = os.environ.get('SLACK_OUATH_TOKEN')

    if(queue_url == None or oauth_token == None):
        print("Missing required environment variables. Exiting...")
        sys.exit()

    # dependencies setup
    sqs = SQSClient()
    slack_webClient = WebClient(token=oauth_token)
    slack = SlackService(slack_webClient)
    event = NotifyAlertEvent(sqs, slack)

    # handles SIGINT and SIGTERM signals to stop process
    signal_handler = SignalHandler()

    # start to handle events
    event.handle(received_signal=signal_handler.received_signal,
                 queue_url=queue_url, wait_time_secs=10)
