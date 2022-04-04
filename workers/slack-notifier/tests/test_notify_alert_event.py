import datetime
import pytest
import boto3
from moto import mock_sqs
from unittest.mock import Mock, patch
from events.notify_alert_event import NotifyAlertEvent
from services.slack_service import SlackService


@pytest.fixture
@mock_sqs
def mock_sqs_client():
    client = boto3.client("sqs")

    return client


@pytest.fixture
def mock_slack_service():
    return Mock(spec=SlackService)


@pytest.fixture
def mock_event(mock_sqs_client, mock_slack_service):
    return NotifyAlertEvent(mock_sqs_client, mock_slack_service)


@mock_sqs
def test_delivery_message_succesfully(mock_slack_service, mock_event):
    sqs = boto3.client("sqs")
    res = sqs.create_queue(QueueName='test_queue')

    queue_url = res["QueueUrl"]
    sqs.send_message(
        QueueUrl=queue_url,
        MessageBody=("{\"channel\":\"teste\",\"message\":\"teste\"}")
    )

    mock_slack_service.post_message.return_value = {'text': 'ok'}

    mock_event.delivery_incoming_messages(
        queue_url=queue_url, wait_time_secs=10)


@mock_sqs
def test_delivery_message_error_handling(mock_slack_service, mock_event):
    sqs = boto3.client("sqs")
    res = sqs.create_queue(QueueName='test_queue')

    queue_url = res["QueueUrl"]
    sqs.send_message(
        QueueUrl=queue_url,
        MessageBody=("{\"channel\":\"teste\",\"message\":\"teste\"}")
    )
    mock_slack_service.post_message.side_effect = Exception()
    with pytest.raises(Exception):
        mock_event.delivery_incoming_messages(
            queue_url=queue_url, wait_time_secs=10)


def test_handle_events(mock_event):
    with patch('events.notify_alert_event.NotifyAlertEvent.delivery_incoming_messages'):
        mock_event.handle(received_signal=True,
                          queue_url='teste_queue', wait_time_secs=3)
