from typing import Any
from abc import ABCMeta, abstractmethod
from dataclasses import dataclass
from slack_sdk import WebClient
from slack_sdk.errors import SlackApiError


@dataclass
class PostMessagePayload:
    channel: str
    text: str


class SlackServiceInterface:
    __metaclass__ = ABCMeta

    @abstractmethod
    async def post_message(self, message: PostMessagePayload) -> Any: pass


class SlackService(SlackServiceInterface):
    def __init__(self, client: WebClient)  -> None:
        self.web_client = client

    async def post_message(self, message: PostMessagePayload):
        try:
            response = self.web_client.chat_postMessage(
                channel=message.channel, text=message.text)

            return response["message"]
        except SlackApiError as e:
            print("There was an error: {}".format(e.response['error']))
            raise
