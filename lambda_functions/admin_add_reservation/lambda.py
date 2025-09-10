import boto3
import json
import hashlib

# Initialize DynamoDB resource
dynamodb = boto3.resource('dynamodb')
venue_info = dynamodb.Table('venue_info')

def lambda_handler(event, context):
    method = event['requestContext']['http']['method']
    path = event['rawPath']
    stage = event['requestContext']['stage']
    path_without_stage = path[len(stage) + 1:] if path.startswith(f"/{stage}/") else path
    queryParams = event.get("queryStringParameters", {})
    print(event)
    
    try:
        if path_without_stage == "/admin/reservations":
            if method == "POST":
                reservationLimitHandler(json.loads(event['body']))
                return {
                    'statusCode': 200,
                    'body': json.dumps({'message': 'Reservation processed successfully'})
                }
    except Exception as error:
        return {
            'statusCode': 500,
            'body': json.dumps({'error': str(error)})
        }


def reservationLimitHandler(message):
    venue = message['venue']
    
    requestId = hashlib.md5((str(message['studentId']) + message['date'] + str(message['time']) + message['venue']).encode()).hexdigest()
    
    response = venue_info.query(
        KeyConditionExpression= boto3.dynamodb.conditions.Key('venue').eq(venue)
    )
    
    items = response.get('Items', [])
    if len(items) == 0:
        return "auto"
    else:
        item = items[0]
        if(not checkDuplicateReservation(message)):
            raise Exception("Duplicate reservation")
        putCurrentReservation(message, requestId)
    
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
            "companion": message['companion']
            "date": year_month,
        }
        current_reservation.put_item(Item=data)
        
    except Exception as e:
        return {
            "statusCode": 500,
            "body": json.dumps({"error": str(e)}),
        }

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
    