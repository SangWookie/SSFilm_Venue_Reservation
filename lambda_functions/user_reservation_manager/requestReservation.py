import boto3
import json

QUEUE_URL = "https://sqs.ap-northeast-2.amazonaws.com/796973485724/reservation-queue.fifo"

def sqs_message_sender(content, message_group):
    sqs_client = boto3.client('sqs', region_name='ap-northeast-2')
    
    response = sqs_client.send_message(
        QueueUrl=QUEUE_URL,
        MessageBody=json.dumps(content),
        MessageGroupId=message_group
    )
    return response

def request_reservation(body):
    try:
        result = sqs_message_sender(body, body['venue'])
        return result
    except Exception as error:
        print("예약 요청 실패:", error)
        raise error