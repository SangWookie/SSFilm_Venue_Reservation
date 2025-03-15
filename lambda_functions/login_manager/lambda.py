import json
import jwt
import datetime

TOKEN_SECRET = "asdfqwer1234"

def lambda_handler(event, context):
    body = event.get("body", {})
    
    try:
        if(body.get("username") == "admin" and body.get("password") == "admin"):
            
            token = issue_access_token(body.get("username"))
            
            return {
                'statusCode': 200,
                'body': json.dumps({'token': token}),
                'headers': {
                    'auth_token': token
                }
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
    payload["iat"] = datetime.datetime.now()
    payload["exp"] = datetime.datetime.now() + datetime.timedelta(minutes=1)
    
    token = jwt.encode(payload, TOKEN_SECRET, algorithm="HS256")
    
    return token

def verify_access_token(token):
    try:
        payload = jwt.decode(token, TOKEN_SECRET, algorithms=["HS256"])
        return payload
    except Exception as error:
        return None