import pytest
import boto3
from moto import mock_aws
from unavailablePeriods import get_unavailable_periods

@pytest.fixture
def dynamodb_setup():
    with mock_aws():
        dynamodb = boto3.resource('dynamodb', region_name='ap-northeast-2')
        reservation_limit = dynamodb.create_table(
            TableName = 'reservation_limit',
            KeySchema = [
                {
                    'AttributeName': 'venueDate',
                    'KeyType': 'HASH'
                }
            ],
            AttributeDefinitions = [
                {
                    'AttributeName': 'venueDate',
                    'AttributeType': 'S'
                }
            ],
            ProvisionedThroughput = {
                "ReadCapacityUnits": 5,
                "WriteCapacityUnits": 5
            }
        )
        
        yield dynamodb
        
def test_get_unavailable_periods(dynamodb_setup):
    reservation_limit = dynamodb_setup.Table('reservation_limit')
    
    reservation_limit.put_item(
        Item={
            "venueDate": "2021-01-01#venue1",
            "limits": [
                {
                    "time": [12, 13],
                    "message": "lunch break"
                },
                {
                    "time": [9, 10, 11],
                    "message": "Closed for maintenance"
                }
            ]
        }
    )
    
    queryParams = {
        "venue": "venue1",
        "date": "2021-01-01"
    }
    
    response = get_unavailable_periods(queryParams, reservation_limit)
    
    assert response == [
        {
            "time": [12, 13],
            "message": "lunch break"
        },
        {
            "time": [9, 10, 11],
            "message": "Closed for maintenance"
        }
    ]
    
def test_empty_response(dynamodb_setup):
    reservation_limit = dynamodb_setup.Table('reservation_limit')
    queryParams = {
        "venue": "venue1",
        "date": "2021-01-01"
    }
    
    response = get_unavailable_periods(queryParams, reservation_limit)
    
    assert response == []