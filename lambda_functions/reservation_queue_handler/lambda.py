import json
import boto3
from reservationLimitHandler import reservationLimitHandler

def lambda_handler(event, context):
    # event 객체에서 SQS 메시지를 추출
    for record in event['Records']:
        # SQS 메시지의 Body를 출력
        body = record['body']
        print(f"Received message: {body}")

        requestId = record['md5OfBody']
        
        message = json.loads(body)
        reservationLimitHandler(message, requestId)
        
        print(f"Processed message: {message}")

    return {
        'statusCode': 200,
        'body': json.dumps(message)
    }