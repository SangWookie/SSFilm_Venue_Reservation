import boto3

dynamodb = boto3.resource('dynamodb', region_name='ap-northeast-2')

def get_unavailable_periods(queryParams, unavailable_periods=None):
    if unavailable_periods is None:
        unavailable_periods = dynamodb.Table('unavailable_periods')
        
    # parameters validation
    if "venue" not in queryParams:
        raise ValueError("The 'venue' attribute is required in queryParams")
    if "date" not in queryParams:
        raise ValueError("The 'date' attribute is required in queryParams")
    
    venue = queryParams['venue']
    date = queryParams['date']
    
    # execute query
    result = unavailable_periods.query(
        KeyConditionExpression= boto3.dynamodb.conditions.Key('venueDate').eq(f"{date}#{venue}")
    )
    
    items = result.get('Items', [])
    
    response = []
    
    for item in items:
        limits = item.get("limits", [])
        time = []
        for limit in limits:
            time = [int(t) for t in limit.get("time", [])]
            message = limit.get("message", "")
            
            response.append({
                "time": time,
                "message": message
            })
    
    return response