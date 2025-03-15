import json
import jwt
import datetime

TOKEN_SECRET = "asdfqwer1234"

def lambda_handler(event, context):
    auth = {
        "isAuthorized": True,
        "context": {
            "username": "admin",
        }
    }
    
    try:
        verify_access_token(event['auth_token'])
        return auth
    
    except Exception as error:
        auth["isAuthorized"] = False
        auth["error"] = str(error)
        return auth
        
        
def verify_access_token(token):
    try:
        payload = jwt.decode(token, TOKEN_SECRET, algorithms=["HS256"])
        return payload
    except Exception as error:
        return None