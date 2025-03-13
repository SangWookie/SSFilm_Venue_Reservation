import json
from requestReservation import request_reservation
from reservationQuery import current_reservation_query
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
                result = current_reservation_query(queryParams)
                unavailable_periods = get_unavailable_periods(queryParams)
                result['unavailable_periods'] = unavailable_periods
                return {
                    'statusCode': 200,
                    'body': json.dumps(result)
                }
            elif method == "POST":
                body = json.loads(event['body'])
                result = request_reservation(body)
                return {
                    'statusCode': 200,
                    'body': json.dumps({'result': result})
                }
        elif path_without_stage == "/reservations/unavailable-periods":
            if method == "GET":
                # query db
                pass
        elif path_without_stage == "/reservations/auto-approval-periods":
            if method == "GET":
                # query db
                pass

        return {
            'statusCode': 200,
            'body': json.dumps({'path': path, 'method': method})
        }
    except Exception as error:
        return {
            'statusCode': 500,
            'body': json.dumps({'error': str(error)})
        }