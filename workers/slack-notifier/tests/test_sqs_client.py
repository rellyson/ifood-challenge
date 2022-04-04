from unittest.mock import patch
import pytest
from aws.sqs_client import SQSClient


@pytest.fixture
def environ_mock(monkeypatch):
    monkeypatch.setenv('AWS_REGION', 'us-east-1')
    monkeypatch.setenv('AWS_ENDPOINT', 'http://test.com')
    monkeypatch.setenv('AWS_ACCESS_KEY_ID', 'test_value')
    monkeypatch.setenv('AWS_SECRET_ACCESS_KEY', 'test_value')


def test_missing_env_vars_validation():
    with pytest.raises(SystemExit) as pytest_wrapped_e:
        _ = SQSClient()
    assert pytest_wrapped_e.type == SystemExit


def test_create_sqs_client(environ_mock):
    client = SQSClient()
    assert client is not None
