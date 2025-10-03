import json
import boto3
import jwt
import datetime
import os
from argon2 import PasswordHasher

ddb = boto3.resource('dynamodb', region_name='ap-northeast-2')
table = ddb.Table('ssfilm_admin_credential')

TOKEN_SECRET = os.environ["token_key"]

ph = PasswordHasher(
    time_cost=1,
    memory_cost=5120,
    parallelism=2
)

def lambda_handler(event, context):
    staged_path = event['rawPath']
    body = json.loads(event["body"])
    stage = event['requestContext']['stage']
    path = staged_path[len(stage) + 1:] if staged_path.startswith(f"/{stage}/") else staged_path

    if path == '/login':
        return try_login(body)
    elif path == '/login/password':
        return change_password(body)
    else:
        return {
            'statusCode': 404,
            'body': json.dumps({'error': 'path not found'})
        }

def try_login(body: dict):
    try:
        response = table.get_item(
            Key={"user_id": body["username"]}
        )
        if "Item" not in response: # 사용자가 존재하지 않음
            raise Exception(f"No user found with username {body['username']}")
        else:
            password = response["Item"]["password"]

        if ph.verify(password, body["password"]):
            token = issue_access_token(body["username"])
            return {
                'statusCode': 200,
                'body': json.dumps({'token': token}),
                'headers': {
                    'auth_token': token
                }
            }

        else:
            return {
                'statusCode': 401,
                'body': json.dumps({'error': 'Invalid credentials'})
            }
    except Exception as error:
        return {
            'statusCode': 500,
            'body': json.dumps({'error': str(error)})
        }

def change_password(body: dict):
    try:
        username = body["username"]
        old_password = body["old_password"]
        new_password = body["new_password"]

        # 사용자 정보 조회
        response = table.get_item(
            Key={"user_id": username}
        )
        if "Item" not in response:
            return {
                'statusCode': 404,
                'body': json.dumps({'error': f"No user found with username {username}"})
            }

        stored_password = response["Item"]["password"]

        # 기존 비밀번호 검증
        if not ph.verify(stored_password, old_password):
            return {
                'statusCode': 401,
                'body': json.dumps({'error': 'Old password is incorrect'})
            }

        # 새 비밀번호 해시 생성
        hashed_new_password = ph.hash(new_password)

        # DynamoDB에 업데이트
        table.update_item(
            Key={"user_id": username},
            UpdateExpression="SET password = :p",
            ExpressionAttributeValues={":p": hashed_new_password}
        )

        return {
            'statusCode': 200,
            'body': json.dumps({'message': 'Password updated successfully'})
        }

    except Exception as error:
        return {
            'statusCode': 500,
            'body': json.dumps({'error': str(error)})
        }


def issue_access_token(username):
    payload = {
        "username": username,
        "exp": datetime.datetime.now(datetime.timezone.utc) + datetime.timedelta(hours=1)
    }

    token = jwt.encode(payload, TOKEN_SECRET, algorithm="HS256")
    
    return token
