import json
import jwt
import datetime
import os
import bcrypt

TOKEN_SECRET = os.environ["token_key"]
USERNAME = os.environ["username"]
PASSWORD = os.environ["password"]

def lambda_handler(event, context):
    body = json.loads(event["body"])
    
    try:
        if(body["username"] == USERNAME and verify_password(body["password"])):
            
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

def issue_access_token(username):
    payload = {
        "username": username
    }
    payload["exp"] = datetime.datetime.now(datetime.timezone.utc) + datetime.timedelta(minutes=10)
    
    token = jwt.encode(payload, TOKEN_SECRET, algorithm="HS256")
    
    return token

def verify_password(password):
    return bcrypt.checkpw(password.encode("utf-8"), PASSWORD.encode("utf-8"))