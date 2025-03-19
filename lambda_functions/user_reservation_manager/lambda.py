import json
from requestReservation import request_reservation
import reservationQuery as r
from unavailablePeriods import get_unavailable_periods

def lambda_handler(event, context):
    method = event['requestContext']['http']['method']
    path = event['rawPath']
    stage = event['requestContext']['stage']
    path_without_stage = path[len(stage) + 1:] if path.startswith(f"/{stage}/") else path
    queryParams = event.get("queryStringParameters", {})

    try:
        if path_without_stage == "/reservations":
            if method == "GET":
                venueResult = r.get_venue_info()
                result = {
                    'date': queryParams.get('date', ''),
                    'venues': venueResult
                }
                for v in result['venues']:
                    venue = v['venue']
                    queryParams['venue'] = venue
                    v['reservations'] = r.current_reservation_query(queryParams)
                    v['unavailable_periods'] = get_unavailable_periods(queryParams)
                return {
                    'statusCode': 200,
                    'body': json.dumps(result)
                }
            elif method == "POST":
                body = json.loads(event['body'])
                result = request_reservation(body)
                return {
                    'statusCode': 200,
                    'body': json.dumps({'reservationId': result['MD5OfMessageBody']})
                }
        elif path_without_stage == "/reservations/check":
            if method == "GET":
                if(r.check_reservation(queryParams) is None):
                    return {
                        'statusCode': 404,
                        'body': 'Reservation not found'
                    }
                else:
                    return {
                        'statusCode': 200,
                        'body': "reservation successful"
                    }
                

        return {
            'statusCode': 200,
            'body': json.dumps({'path': path, 'method': method})
        }
    except Exception as error:
        return {
            'statusCode': 500,
            'body': json.dumps({'error': str(error)})
        }