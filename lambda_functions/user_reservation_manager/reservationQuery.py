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
    
    response = {
        "date": date,
        "venue": venue,
        "reservations": [],
    }
    
    for item in items:
        converted_time = [int(x) for x in item.get("time", [])]
        
        response["reservations"].append({
            "name": item.get("name", ""),
            "time": converted_time,
            "purpose": item.get("purpose", ""),
            "category": item.get("category", ""),
        })
    
    return response