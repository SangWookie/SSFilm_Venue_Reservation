import pytest
import boto3
from moto import mock_aws
from reservationLimitHandler import reservationLimitHandler, putCurrentReservation, putPendingReservation

@pytest.fixture
def dynamodb_setup():
    with mock_aws():
        dynamodb = boto3.resource('dynamodb', region_name='ap-northeast-2')
        venue_info = dynamodb.create_table(
            TableName = 'venue_info',
            KeySchema = [
                {
                    'AttributeName': 'venue',
                    'KeyType': 'HASH'
                }
            ],
            AttributeDefinitions = [
                {
                    'AttributeName': 'venue',
                    'AttributeType': 'S'
                }
            ],
            ProvisionedThroughput = {
                "ReadCapacityUnits": 5,
                "WriteCapacityUnits": 5
            }
        )
        
        current_reservation = dynamodb.create_table(
            TableName='current_reservation',
            KeySchema=[
                {
                    'AttributeName': 'reservationId',  # Hash key
                    'KeyType': 'HASH'  # Partition key
                }
            ],
            AttributeDefinitions=[
                {
                    'AttributeName': 'reservationId',
                    'AttributeType': 'S'  # String type
                },
                {
                    'AttributeName': 'venueDate',
                    'AttributeType': 'S'  # String type for GSI
                }
            ],
            ProvisionedThroughput={
                'ReadCapacityUnits': 5,
                'WriteCapacityUnits': 5
            },
            GlobalSecondaryIndexes=[
                {
                    'IndexName': 'venueDate-index',  # GSI name
                    'KeySchema': [
                        {
                            'AttributeName': 'venueDate',
                            'KeyType': 'HASH'  # GSI hash key
                        }
                    ],
                    'Projection': {
                        'ProjectionType': 'ALL'  # GSI에 저장될 모든 속성
                    },
                    'ProvisionedThroughput': {
                        'ReadCapacityUnits': 5,
                        'WriteCapacityUnits': 5
                    }
                }
            ]
        )
        
        pendingReservation = dynamodb.create_table(
            TableName='pending_reservation',
            KeySchema=[
                {
                    'AttributeName': 'requestId',
                    'KeyType': 'HASH'
                }
            ],
            AttributeDefinitions=[
                {
                    'AttributeName': 'requestId',
                    'AttributeType': 'S'
                }
            ],
            ProvisionedThroughput={
                'ReadCapacityUnits': 5,
                'WriteCapacityUnits': 5
            }
        )

        dynamodb.Table('venue_info').put_item(Item={
            'venue': 'studio',
            'venueKor': '스튜디오',
            'allowPolicy': 'auto'
        })
        dynamodb.Table('venue_info').put_item(Item={
            'venue': 'lounge',
            'venueKor': '라운지',
            'allowPolicy': 'manual'
        })
        yield dynamodb


def test_venue_info_insert(dynamodb_setup):    
    venue_info = dynamodb_setup.Table('venue_info')
    assert venue_info.query(KeyConditionExpression=boto3.dynamodb.conditions.Key('venue').eq('studio'))['Items'][0]['allowPolicy'] == 'auto'
    

def test_reservation_auto_policy(dynamodb_setup):
    """자동 승인(`auto`)일 경우 current_reservation 테이블에 추가되는지 테스트"""

    message = {
        "venue": "studio",
        "date": "2025-03-12",
        "time": [10,11],
        "studentId": 20201728,
        "name": "John Doe",
        "email": "john@example.com",
        "category": "study",
        "purpose": "group study",
        "companion": "me and you",
    }
    
    reservationLimitHandler(message, "asdfweif23f2")

    current_reservation = dynamodb_setup.Table("current_reservation")
    response = current_reservation.scan()
    assert len(response["Items"]) == 1
    assert response["Items"][0]["studentId"] == 20201728
    assert response["Items"][0]["venueDate"] == "2025-03-12#studio"
    assert response["Items"][0]["time"] == [10, 11]


def test_reservation_manual_policy(dynamodb_setup):
    """수동 승인(`manual`)일 경우 pending_reservation 테이블에 추가되는지 테스트"""
    message = {
        "venue": "lounge",
        "date": "2025-03-12",
        "time": [10,11],
        "studentId": 20201728,
        "name": "John Doe",
        "email": "john@example.com",
        "category": "study",
        "purpose": "group study",
        "companion": "me and you",
    }
    
    reservationLimitHandler(message, "req-002")

    pending_reservation = dynamodb_setup.Table("pending_reservation")
    response = pending_reservation.scan()
    assert len(response["Items"]) == 1
    assert response["Items"][0]["studentId"] == 20201728
    assert response["Items"][0]["venueDate"] == "2025-03-12#lounge"
    assert response["Items"][0]["time"] == [10, 11]


def test_6hours_limit(dynamodb_setup):
    """하루에 6시간 이상 예약할 수 없는지 테스트"""
    message = {
        "venue": "studio",
        "date": "2025-03-12",
        "time": [10,11,12,13,14],
        "studentId": 20201728,
        "name": "John Doe",
        "email": "john@example.com",
        "category": "study",
        "purpose": "group study",
        "companion": "me and you",
    }
    message2 = {
        "venue": "studio",
        "date": "2025-03-12",
        "time": [15, 16],
        "studentId": 20201728,
        "name": "John Doe",
        "email": "john@example.com",
        "category": "study",
        "purpose": "group study",
        "companion": "me and you",
    }
    reservationLimitHandler(message, "req-001")
    with pytest.raises(Exception):
        reservationLimitHandler(message2, "req-002")

def test_multiple_reservations(dynamodb_setup):
    """하루에 6시간 미만 예약할 수 있는지지 테스트"""
    message = {
        "venue": "studio",
        "date": "2025-03-12",
        "time": [10,11,12,13,14],
        "studentId": 20201728,
        "name": "John Doe",
        "email": "john@example.com",
        "category": "study",
        "purpose": "group study",
        "companion": "me and you",
    }
    message2 = {
        "venue": "studio",
        "date": "2025-03-12",
        "time": [15],
        "studentId": 20201728,
        "name": "John Doe",
        "email": "john@example.com",
        "category": "study",
        "purpose": "group study",
        "companion": "me and you",
    }
    reservationLimitHandler(message, "req-001")
    reservationLimitHandler(message2, "req-002")
    
    current_reservation = dynamodb_setup.Table("current_reservation")
    response = current_reservation.scan()
    items = response.get("Items", [])
    assert len(items) == 2

def test_duplicate_reservation(dynamodb_setup):
    """중복 시간대에 예약이 불가능한지 테스트"""
    message = {
        "venue": "studio",
        "date": "2025-03-12",
        "time": [11,12,13,14],
        "studentId": 20201728,
        "name": "John Doe",
        "email": "john@example.com",
        "category": "study",
        "purpose": "group study",
        "companion": "me and you",
    }
    reservationLimitHandler(message, "req-001")
    
    message2 = {
        "venue": "studio",
        "date": "2025-03-12",
        "time": [14,15],
        "studentId": 20201728,
        "name": "John Doe",
        "email": "john@example.com",
        "category": "study",
        "purpose": "group study",
        "companion": "me and you",
    }
    
    with pytest.raises(Exception):
        reservationLimitHandler(message2, "req-002")