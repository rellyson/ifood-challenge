import asyncio
import pytest
from unittest.mock import patch
from slack_sdk.errors import SlackApiError
from slack_sdk import WebClient
from services.slack_service import SlackService, PostMessagePayload


@pytest.fixture
def mock_service():
    return SlackService(WebClient())


@patch('services.slack_service.WebClient.chat_postMessage')
def test_post_message_succesful_response(mock_client, mock_service):

    payload = PostMessagePayload(channel="test", text="test", attachments=[])

    slack_response_mock = {
        'message': {
            'teste': 'teste'
        }
    }

    mock_client.return_value = slack_response_mock

    res = asyncio.run(mock_service.post_message(payload))

    assert res == slack_response_mock['message']


def test_post_message_exception_handling(mock_service):
    with patch('services.slack_service.WebClient.chat_postMessage', side_effect=SlackApiError(
            "test error", {'error': 'test error'})):
        with pytest.raises(SlackApiError):
            payload = PostMessagePayload(channel="test", text="test", attachments=[])
            asyncio.run(mock_service.post_message(payload))
