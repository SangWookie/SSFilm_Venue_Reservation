import json
import jwt
import os

TOKEN_SECRET = os.environ["token_key"]

def lambda_handler(event, context):
    auth = {
        "isAuthorized": True,
        "context": {
            "username": "admin",
        }
    }
    print(event)
    try:
        token = event['identitySource'][0]
        verify_access_token(token)
        return auth
    except Exception as error:
        auth["isAuthorized"] = False
        print(error)
        return auth
        
        
def verify_access_token(token):
    payload = jwt.decode(token, TOKEN_SECRET, algorithms=["HS256"])
    return payload
