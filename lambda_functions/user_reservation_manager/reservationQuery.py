import boto3

dynamodb = boto3.resource('dynamodb', region_name='ap-northeast-2')

def current_reservation_query(queryParams, current_reservation=None):
    if current_reservation is None:
        current_reservation = dynamodb.Table('current_reservation')
        
    # parameters validation
    if "venue" not in queryParams:
        raise ValueError("The 'venue' attribute is required in queryParams")
    if "date" not in queryParams:
        raise ValueError("The 'date' attribute is required in queryParams")
    
    venue = queryParams['venue']
    date = queryParams['date']
    
    # execute query
    result = current_reservation.query(
        IndexName='venueDate-index',
        KeyConditionExpression= "venueDate = :venueDate",
        ExpressionAttributeValues= {":venueDate": f"{date}#{venue}"}
    )
    
    items = result.get('Items', [])
    
    response = []
    
    for item in items:
        converted_time = [int(x) for x in item.get("time", [])]
        
        response.append({
            "name": item.get("name", ""),
            "time": converted_time,
            "category": item.get("category", ""),
            "purpose": item.get("purpose", ""),
        })
    
    return response

def check_reservation(queryParams):
    if("reservationId" not in queryParams):
        raise ValueError("The 'reservationId' attribute is required in queryParams")
    
    reservation_id = queryParams['reservationId']
    current_reservation = dynamodb.Table('current_reservation')
    result = current_reservation.query(
        KeyConditionExpression= "reservationId = :reservationId",
        ExpressionAttributeValues= {":reservationId": reservation_id}
    )
    
    Items = result.get('Items', [])
    
    if len(Items) == 0:
        return None
    else:
        return Items[0]
    

venue_info = dynamodb.Table('venue_info')

def get_venue_info():
    venueResult = venue_info.scan()
    items = venueResult.get('Items', [])
    venues = []
    for item in items:
        venues.append({
            "venue": item.get("venue", ""),
            "venueKor": item.get("venueKor", ""),
            "approval_mode": item.get("allowPolicy", ""),
        })
    return venues
