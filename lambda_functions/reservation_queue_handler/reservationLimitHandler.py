import boto3
import boto3.dynamodb
import boto3.dynamodb.conditions
import json

# Initialize DynamoDB resource
dynamodb = boto3.resource('dynamodb')
venue_info = dynamodb.Table('venue_info')

def reservationLimitHandler(message, requestId):
    venue = message['venue']
    
    response = venue_info.query(
        KeyConditionExpression= boto3.dynamodb.conditions.Key('venue').eq(venue)
    )
    
    items = response.get('Items', [])
    if len(items) == 0:
        return "auto"
    else:
        item = items[0]
        if(item['allowPolicy'] == 'auto'):
            if(not checkDailyReservationLimit(message)):
                raise Exception("Daily reservation limit exceeded")
            if(not checkDuplicateReservation(message)):
                raise Exception("Duplicate reservation")
            putCurrentReservation(message, requestId)
        elif(item['allowPolicy'] == 'manual'):
            putPendingReservation(message, requestId)
    
current_reservation = dynamodb.Table('current_reservation')

def putCurrentReservation(message, requestId):
    full_date = message['date']
    year_month = full_date[:7]
    try:
        data = {
            "reservationId": requestId,
            "venueDate": full_date + '#' + message['venue'],
            "time": message['time'],
            "studentId": message['studentId'],
            "name": message['name'],
            "email": message['email'],
            "category": message['category'],
            "purpose": message['purpose'],
            "companion": message['companion'],
            "date": year_month
        }
        current_reservation.put_item(Item=data)
        
    except Exception as e:
        print(e)
        return {
            "statusCode": 500,
            "body": json.dumps({"error": str(e)}),
        }

pending_reservation = dynamodb.Table('pending_reservation')

def putPendingReservation(message, requestId):
    try:
        data = {
            "requestId": requestId,
            "venueDate": message['date'] + '#' + message['venue'],
            "time": message['time'],
            "studentId": message['studentId'],
            "name": message['name'],
            "email": message['email'],
            "category": message['category'],
            "purpose": message['purpose'],
            "companion": message['companion']
        }
        pending_reservation.put_item(Item=data)
        
    except Exception as e:
        print(e)
        return {
            "statusCode": 500,
            "body": json.dumps({"error": str(e)}),
        }
        
def checkDailyReservationLimit(message):
    date = message['date']
    
    response = current_reservation.scan(
        FilterExpression= boto3.dynamodb.conditions.Attr('venueDate').begins_with(date)&
            boto3.dynamodb.conditions.Attr('studentId').eq(message['studentId'])
    )
    
    items = response.get('Items', [])
    
    if(items == []):
        return True
    
    timeCnt = len(message['time'])
    for item in items:
        timeCnt += len(item['time'])
    
    if(timeCnt > 6):
        print("Daily reservation limit exceeded")
        return False
    else:
        return True

def checkDuplicateReservation(message):
    venue = message['venue']
    date = message['date']
    
    response = current_reservation.scan(
        FilterExpression= boto3.dynamodb.conditions.Attr('venueDate').eq(date + '#' + venue)
    )
    
    items = response.get('Items', [])
    
    if(items == []):
        return True
    
    occupied = []
    for item in items:
        occupied.extend(item['time'])
        print("occupied: ", occupied)
    
    if(bool(set(occupied) & set(message['time']))) :
        return False
    else:
        return True
    