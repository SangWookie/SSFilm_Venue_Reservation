package test

import (
	"context"
	"encoding/json"
	"request_manager/handlers"
	"request_manager/mocks"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestManageReservation_Cancel(t *testing.T) {
	// Setup mock DynamoDB client
	mockDDB := &mocks.MockDDBClient{}

	// Mock reservation data
	reservationID := "test-reservation-id"

	// Mock IsItemExist response
	mockDDB.On("GetItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.GetItemInput) bool {
		return *input.TableName == "current_reservation" &&
			input.Key["reservationID"].(*types.AttributeValueMemberS).Value == reservationID
	})).Return(&dynamodb.GetItemOutput{
		Item: map[string]types.AttributeValue{
			"reservationID": &types.AttributeValueMemberS{Value: reservationID},
		},
	}, nil)

	// Mock DeleteItem response
	mockDDB.On("DeleteItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.DeleteItemInput) bool {
		return *input.TableName == "pending_reservation" &&
			input.Key["reservationID"].(*types.AttributeValueMemberS).Value == reservationID
	})).Return(&dynamodb.DeleteItemOutput{}, nil)

	// Create request body
	requestBody := handlers.RequestChangeTimeRequest{
		Key:  reservationID,
		Code: "CANCEL",
	}
	bodyBytes, _ := json.Marshal(requestBody)

	// Create API Gateway request
	ctx := context.Background()
	request := events.APIGatewayV2HTTPRequest{
		Body: string(bodyBytes),
	}

	// Call the handler
	response, err := handlers.ManageReservation(ctx, request, mockDDB)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)

	// Verify all mocks were called
	mockDDB.AssertExpectations(t)
}

func TestManageReservation_Modify(t *testing.T) {
	// Setup mock DynamoDB client
	mockDDB := &mocks.MockDDBClient{}

	// Mock reservation data
	reservationID := "test-reservation-id"

	// Create time values for the test
	timeValues := []types.AttributeValue{
		&types.AttributeValueMemberN{Value: "13"},
		&types.AttributeValueMemberN{Value: "14"},
		&types.AttributeValueMemberN{Value: "15"},
	}

	// Mock IsItemExist response
	mockDDB.On("GetItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.GetItemInput) bool {
		return *input.TableName == "current_reservation" &&
			input.Key["reservationID"].(*types.AttributeValueMemberS).Value == reservationID
	})).Return(&dynamodb.GetItemOutput{
		Item: map[string]types.AttributeValue{
			"reservationID": &types.AttributeValueMemberS{Value: reservationID},
		},
	}, nil)

	// Mock UpdateItem response
	mockDDB.On("UpdateItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.UpdateItemInput) bool {
		return *input.TableName == "current_reservation" &&
			input.Key["reservationID"].(*types.AttributeValueMemberS).Value == reservationID
	})).Return(&dynamodb.UpdateItemOutput{}, nil)

	// Create request body
	requestBody := handlers.RequestChangeTimeRequest{
		Key:        reservationID,
		Code:       "MODIFY",
		ChangeTime: timeValues,
	}
	bodyBytes, _ := json.Marshal(requestBody)

	// Create API Gateway request
	ctx := context.Background()
	request := events.APIGatewayV2HTTPRequest{
		Body: string(bodyBytes),
	}

	// Call the handler
	response, err := handlers.ManageReservation(ctx, request, mockDDB)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)

	// Verify all mocks were called
	mockDDB.AssertExpectations(t)
}

func TestManageReservation_ItemNotFound(t *testing.T) {
	// Setup mock DynamoDB client
	mockDDB := &mocks.MockDDBClient{}

	// Mock reservation data
	reservationID := "non-existent-id"

	// Mock IsItemExist response for non-existent item
	mockDDB.On("GetItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.GetItemInput) bool {
		return *input.TableName == "current_reservation" &&
			input.Key["reservationID"].(*types.AttributeValueMemberS).Value == reservationID
	})).Return(&dynamodb.GetItemOutput{
		Item: nil, // No item found
	}, nil)

	// Create request body
	requestBody := handlers.RequestChangeTimeRequest{
		Key:  reservationID,
		Code: "CANCEL",
	}
	bodyBytes, _ := json.Marshal(requestBody)

	// Create API Gateway request
	ctx := context.Background()
	request := events.APIGatewayV2HTTPRequest{
		Body: string(bodyBytes),
	}

	// Call the handler
	response, err := handlers.ManageReservation(ctx, request, mockDDB)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, "Not found Item", response.Body)

	// Verify all mocks were called
	mockDDB.AssertExpectations(t)
}

func TestManageReservation_InvalidCode(t *testing.T) {
	// Setup mock DynamoDB client
	mockDDB := &mocks.MockDDBClient{}

	// Mock reservation data
	reservationID := "test-reservation-id"

	// Mock IsItemExist response
	mockDDB.On("GetItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.GetItemInput) bool {
		return *input.TableName == "current_reservation" &&
			input.Key["reservationID"].(*types.AttributeValueMemberS).Value == reservationID
	})).Return(&dynamodb.GetItemOutput{
		Item: map[string]types.AttributeValue{
			"reservationID": &types.AttributeValueMemberS{Value: reservationID},
		},
	}, nil)

	// Create request body with invalid code
	requestBody := handlers.RequestChangeTimeRequest{
		Key:  reservationID,
		Code: "INVALID_CODE",
	}
	bodyBytes, _ := json.Marshal(requestBody)

	// Create API Gateway request
	ctx := context.Background()
	request := events.APIGatewayV2HTTPRequest{
		Body: string(bodyBytes),
	}

	// Call the handler
	response, err := handlers.ManageReservation(ctx, request, mockDDB)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)

	// Verify mocks were called
	mockDDB.AssertExpectations(t)
}
