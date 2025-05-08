import boto3
import asyncio
from unavailablePeriods import get_unavailable_periods

dynamodb = boto3.resource('dynamodb', region_name='ap-northeast-2')

async def current_reservation_query(queryParams, current_reservation=None):
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
    sorted_items = sorted(items, key=lambda x: x.get("sortOrder", 999))
    venues = []
    for item in sorted_items:
        venues.append({
            "venue": item.get("venue", ""),
            "venueKor": item.get("venueKor", ""),
            "approval_mode": item.get("allowPolicy", ""),
        })
    return venues

async def getResrv(queryParams):
    venueResult = get_venue_info()  # 여기는 동기 호출이라고 가정
    result = {
        'date': queryParams.get('date', ''),
        'venues': venueResult
    }

    # venue 목록을 비동기적으로 처리
    tasks = [enrich_venue(v, queryParams) for v in result['venues']]
    result['venues'] = await asyncio.gather(*tasks)

    return result

async def enrich_venue(v, base_query):
    venue = v['venue']
    query = base_query.copy()
    query['venue'] = venue

    # 두 개의 I/O 작업을 병렬적으로 실행
    res_task = asyncio.create_task(current_reservation_query(query))
    unavail_task = asyncio.create_task(get_unavailable_periods(query))

    v['reservations'], v['unavailable_periods'] = await asyncio.gather(res_task, unavail_task)
    return v