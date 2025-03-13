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

func TestManagePendingReservation_Accept(t *testing.T) {
	// Setup mock DynamoDB client
	mockDDB := &mocks.MockDDBClient{}

	// Mock reservation data
	requestId := "test-reservation-id"
	pendingReservation := map[string]types.AttributeValue{
		"requestId": &types.AttributeValueMemberS{Value: requestId},
		"category":  &types.AttributeValueMemberS{Value: "수업"},
		"companion": &types.AttributeValueMemberS{Value: "james, andrew"},
		"email":     &types.AttributeValueMemberS{Value: "tester@tester.com"},
		"name":      &types.AttributeValueMemberS{Value: "tester"},
		"purpose":   &types.AttributeValueMemberS{Value: "specific purpose for usage"},
		"studentId": &types.AttributeValueMemberN{Value: "20201728"},
		"time": &types.AttributeValueMemberL{
			Value: []types.AttributeValue{
				&types.AttributeValueMemberN{Value: "10"},
				&types.AttributeValueMemberN{Value: "11"},
				&types.AttributeValueMemberN{Value: "12"},
			},
		},
		"venueDate": &types.AttributeValueMemberS{Value: "2025-03-31#studio"},
	}

	// Mock GetItem response
	mockDDB.On("GetItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.GetItemInput) bool {
		return *input.TableName == "pending_reservation" &&
			input.Key["requestId"].(*types.AttributeValueMemberS).Value == requestId
	})).Return(&dynamodb.GetItemOutput{
		Item: pendingReservation,
	}, nil)

	// Mock PutItem response for AcceptReservation
	mockDDB.On("PutItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.PutItemInput) bool {
		return *input.TableName == "current_reservation"
	})).Return(&dynamodb.PutItemOutput{}, nil)

	// Mock DeleteItem response
	mockDDB.On("DeleteItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.DeleteItemInput) bool {
		return *input.TableName == "pending_reservation" &&
			input.Key["requestId"].(*types.AttributeValueMemberS).Value == requestId
	})).Return(&dynamodb.DeleteItemOutput{}, nil)

	// Create request body
	requestBody := handlers.RequestDeleteType{
		Key:  requestId,
		Code: "ACCEPT",
	}
	bodyBytes, _ := json.Marshal(requestBody)

	// Create API Gateway request
	ctx := context.Background()
	request := events.APIGatewayV2HTTPRequest{
		Body: string(bodyBytes),
	}

	// Call the handler
	response, err := handlers.ManagePendingReservation(ctx, request, mockDDB)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)

	// Verify all mocks were called
	mockDDB.AssertExpectations(t)
}

func TestManagePendingReservation_Deny(t *testing.T) {
	// Setup mock DynamoDB client
	mockDDB := &mocks.MockDDBClient{}

	// Mock reservation data
	requestId := "test-reservation-id"
	pendingReservation := map[string]types.AttributeValue{
		"requestId": &types.AttributeValueMemberS{Value: requestId},
		"category":  &types.AttributeValueMemberS{Value: "수업"},
		"companion": &types.AttributeValueMemberS{Value: "james, andrew"},
		"email":     &types.AttributeValueMemberS{Value: "tester@tester.com"},
		"name":      &types.AttributeValueMemberS{Value: "tester"},
		"purpose":   &types.AttributeValueMemberS{Value: "specific purpose for usage"},
		"studentId": &types.AttributeValueMemberN{Value: "20201728"},
		"time": &types.AttributeValueMemberL{
			Value: []types.AttributeValue{
				&types.AttributeValueMemberN{Value: "10"},
				&types.AttributeValueMemberN{Value: "11"},
				&types.AttributeValueMemberN{Value: "12"},
			},
		},
		"venueDate": &types.AttributeValueMemberS{Value: "2025-03-31#studio"},
	}

	// Mock GetItem response
	mockDDB.On("GetItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.GetItemInput) bool {
		return *input.TableName == "pending_reservation" &&
			input.Key["requestId"].(*types.AttributeValueMemberS).Value == requestId
	})).Return(&dynamodb.GetItemOutput{
		Item: pendingReservation,
	}, nil)

	// Mock DeleteItem response
	mockDDB.On("DeleteItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.DeleteItemInput) bool {
		return *input.TableName == "pending_reservation" &&
			input.Key["requestId"].(*types.AttributeValueMemberS).Value == requestId
	})).Return(&dynamodb.DeleteItemOutput{}, nil)

	// Create request body
	requestBody := handlers.RequestDeleteType{
		Key:  requestId,
		Code: "DENY",
	}
	bodyBytes, _ := json.Marshal(requestBody)

	// Create API Gateway request
	ctx := context.Background()
	request := events.APIGatewayV2HTTPRequest{
		Body: string(bodyBytes),
	}

	// Call the handler
	response, err := handlers.ManagePendingReservation(ctx, request, mockDDB)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)

	// Verify all mocks were called
	mockDDB.AssertExpectations(t)
}

func TestManagePendingReservation_InvalidCode(t *testing.T) {
	// Setup mock DynamoDB client
	mockDDB := &mocks.MockDDBClient{}

	// Mock reservation data
	requestId := "test-reservation-id"
	pendingReservation := map[string]types.AttributeValue{
		"requestId": &types.AttributeValueMemberS{Value: requestId},
	}

	// Mock GetItem response
	mockDDB.On("GetItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.GetItemInput) bool {
		return *input.TableName == "pending_reservation" &&
			input.Key["requestId"].(*types.AttributeValueMemberS).Value == requestId
	})).Return(&dynamodb.GetItemOutput{
		Item: pendingReservation,
	}, nil)

	// Create request body with invalid code
	requestBody := handlers.RequestDeleteType{
		Key:  requestId,
		Code: "INVALID_CODE",
	}
	bodyBytes, _ := json.Marshal(requestBody)

	// Create API Gateway request
	ctx := context.Background()
	request := events.APIGatewayV2HTTPRequest{
		Body: string(bodyBytes),
	}

	// Call the handler
	response, err := handlers.ManagePendingReservation(ctx, request, mockDDB)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, 404, response.StatusCode)

	// Verify mocks were called
	mockDDB.AssertExpectations(t)
}
