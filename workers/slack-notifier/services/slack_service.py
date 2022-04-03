import os
from dataclasses import dataclass
from abc import ABCMeta, abstractmethod
from slack_sdk import WebClient
from slack_sdk.errors import SlackApiError


@dataclass
class MessagePayload:
    channel: str
    text: str


class SlackServiceInterface:
    __metaclass__ = ABCMeta

    @abstractmethod
    async def post_message(self, message: MessagePayload) -> None: pass


class SlackService(SlackServiceInterface):
    def __init__(self) -> None:
        self.web_client = WebClient(token=os.environ.get('SLACK_BOT_TOKEN'))

    async def post_message(self, message: MessagePayload):
        try:
            response = self.web_client.chat_postMessage(
                channel=message.channel, text=message.text)
            assert response["message"]["text"] == message.text

        except SlackApiError as e:
            print("There was an error: {}".format(e.response['error']))
            raise
