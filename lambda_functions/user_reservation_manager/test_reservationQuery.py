import pytest
import boto3
from moto import mock_aws
from reservationQuery import current_reservation_query

@pytest.fixture
def dynamodb_setup():
    with mock_aws():
        dynamodb = boto3.resource('dynamodb', region_name='ap-northeast-2')
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
        
        current_reservation.put_item(Item={
            'reservationId': '1',
            'venueDate': '2025-03-01#studio',
            'time': [10, 11, 12],
            'studentId': 20201728,
            'name': 'name1',
            'email': 'email1',
            'category': 'category1',
            'purpose': 'purpose1',
            'companion': 'companion1'
        })
        current_reservation.put_item(Item={
            'reservationId': '2',
            'venueDate': '2025-03-01#studio',
            'time': [9, 10],
            'studentId': 20201111,
            'name': 'hello',
            'email': 'email1',
            'category': 'category1',
            'purpose': 'purpose1',
            'companion': 'companion1'
        })
        
        yield dynamodb
        
def test_basic_query(dynamodb_setup):
    requestParams = {
        'venue': 'studio',
        'date': '2025-03-01'
    }
    response = current_reservation_query(requestParams, current_reservation=dynamodb_setup.Table('current_reservation'))
    
    assert response['date'] == '2025-03-01'
    assert response['venue'] == 'studio'
    assert response['reservations'][0]['time'] == [10, 11, 12]
    assert len(response['reservations']) == 2