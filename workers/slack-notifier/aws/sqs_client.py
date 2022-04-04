import sys
import boto3
import os


def SQSClient() -> boto3.client:
    region = os.environ.get('AWS_REGION')
    endpoint = os.environ.get('AWS_ENDPOINT')
    access_id = os.environ.get('AWS_ACCESS_KEY_ID')
    access_key = os.environ.get('AWS_SECRET_ACCESS_KEY')

    if(endpoint == None or region == None or access_id == None or access_key == None):
        print("Missing AWS required environment variables")
        sys.exit()
    return boto3.client("sqs", region_name=region,
                        endpoint_url=endpoint,
                        aws_access_key_id=access_id,
                        aws_secret_access_key=access_key)
