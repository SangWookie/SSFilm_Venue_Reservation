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
	mockSMTP := &mocks.MockSendEmail{}

	// Mock reservation data
	reservationID := "test-reservation-id"

	// Mock IsItemExist response with email
	mockDDB.On("GetItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.GetItemInput) bool {
		return *input.TableName == "current_reservation" &&
			input.Key["reservationId"].(*types.AttributeValueMemberS).Value == reservationID
	})).Return(&dynamodb.GetItemOutput{
		Item: map[string]types.AttributeValue{
			"reservationId": &types.AttributeValueMemberS{Value: reservationID},
			"email":         &types.AttributeValueMemberS{Value: "test@example.com"},
			"venueDate":     &types.AttributeValueMemberS{Value: "2025-03-03#studio"},
			"name":          &types.AttributeValueMemberS{Value: "test"},
			"category":      &types.AttributeValueMemberS{Value: "test"},
			"time": &types.AttributeValueMemberL{
				Value: []types.AttributeValue{
					&types.AttributeValueMemberN{Value: "10"},
					&types.AttributeValueMemberN{Value: "11"},
					&types.AttributeValueMemberN{Value: "12"},
				},
			},
		},
	}, nil)

	// Mock DeleteItem response
	mockDDB.On("DeleteItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.DeleteItemInput) bool {
		return *input.TableName == "current_reservation" &&
			input.Key["reservationId"].(*types.AttributeValueMemberS).Value == reservationID
	})).Return(&dynamodb.DeleteItemOutput{}, nil)

	mockSMTP.On("SendEmailWithGoogle", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Create request body
	requestBody := handlers.RequestChangeRequest{
		Key:    reservationID,
		Code:   "CANCEL",
		Reason: "취소 사유",
	}
	bodyBytes, _ := json.Marshal(requestBody)

	// Create API Gateway request
	ctx := context.Background()
	request := events.APIGatewayV2HTTPRequest{
		Body: string(bodyBytes),
	}

	// Create handler parameters
	params := handlers.RouterHandlerParameters{
		Ctx:        ctx,
		Request:    request,
		DdbClient:  mockDDB,
		SmtpClient: mockSMTP,
	}

	// Call the handler
	response, err := handlers.ManageReservation(params)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)

	// Verify all mocks were called
	mockDDB.AssertExpectations(t)
	mockSMTP.AssertExpectations(t)
}

func TestManageReservation_Modify(t *testing.T) {
	// Setup mock DynamoDB client
	mockDDB := &mocks.MockDDBClient{}
	mockSMTP := &mocks.MockSendEmail{}

	// Mock reservation data
	reservationID := "test-reservation-id"

	// Create time values for the test
	timeValues := []int{13, 14, 15}

	// Mock IsItemExist response with email
	mockDDB.On("GetItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.GetItemInput) bool {
		return *input.TableName == "current_reservation" &&
			input.Key["reservationId"].(*types.AttributeValueMemberS).Value == reservationID
	})).Return(&dynamodb.GetItemOutput{
		Item: map[string]types.AttributeValue{
			"reservationId": &types.AttributeValueMemberS{Value: reservationID},
			"email":         &types.AttributeValueMemberS{Value: "test@example.com"},
			"venueDate":     &types.AttributeValueMemberS{Value: "2025-03-03#studio"},
			"name":          &types.AttributeValueMemberS{Value: "test"},
			"category":      &types.AttributeValueMemberS{Value: "test"},
			"time": &types.AttributeValueMemberL{
				Value: []types.AttributeValue{
					&types.AttributeValueMemberN{Value: "10"},
					&types.AttributeValueMemberN{Value: "11"},
					&types.AttributeValueMemberN{Value: "12"},
				},
			},
		},
	}, nil)

	// Mock UpdateItem response
	mockDDB.On("UpdateItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.UpdateItemInput) bool {
		return *input.TableName == "current_reservation" &&
			input.Key["reservationId"].(*types.AttributeValueMemberS).Value == reservationID
	})).Return(&dynamodb.UpdateItemOutput{}, nil)

	mockSMTP.On("SendEmailWithGoogle", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Create request body
	requestBody := handlers.RequestChangeRequest{
		Key:  reservationID,
		Code: "MODIFY",
		ChangeValues: handlers.ChangeValuesType{
			ChangeTime: timeValues,
			Venue:      "studio",
			Date:       "2025-03-03",
		},
		Reason: "변경 사유",
	}
	bodyBytes, _ := json.Marshal(requestBody)

	// Create API Gateway request
	ctx := context.Background()
	request := events.APIGatewayV2HTTPRequest{
		Body: string(bodyBytes),
	}

	// Create handler parameters
	params := handlers.RouterHandlerParameters{
		Ctx:        ctx,
		Request:    request,
		DdbClient:  mockDDB,
		SmtpClient: mockSMTP,
	}

	// Call the handler
	response, err := handlers.ManageReservation(params)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)

	// Verify all mocks were called
	mockDDB.AssertExpectations(t)
	mockSMTP.AssertExpectations(t)
}

func TestManageReservation_CancelEmailError(t *testing.T) {
	// Setup mock DynamoDB client
	mockDDB := &mocks.MockDDBClient{}
	mockSMTP := &mocks.MockSendEmail{}

	// Mock reservation data
	reservationID := "test-reservation-id"

	// Mock IsItemExist response with email
	mockDDB.On("GetItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.GetItemInput) bool {
		return *input.TableName == "current_reservation" &&
			input.Key["reservationId"].(*types.AttributeValueMemberS).Value == reservationID
	})).Return(&dynamodb.GetItemOutput{
		Item: map[string]types.AttributeValue{
			"reservationId": &types.AttributeValueMemberS{Value: reservationID},
			"email":         &types.AttributeValueMemberS{Value: "test@example.com"},
			"venueDate":     &types.AttributeValueMemberS{Value: "2025-03-03#studio"},
			"name":          &types.AttributeValueMemberS{Value: "test"},
			"category":      &types.AttributeValueMemberS{Value: "test"},
			"time": &types.AttributeValueMemberL{
				Value: []types.AttributeValue{
					&types.AttributeValueMemberN{Value: "10"},
					&types.AttributeValueMemberN{Value: "11"},
					&types.AttributeValueMemberN{Value: "12"},
				},
			},
		},
	}, nil)

	// Mock DeleteItem response
	mockDDB.On("DeleteItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.DeleteItemInput) bool {
		return *input.TableName == "current_reservation" &&
			input.Key["reservationId"].(*types.AttributeValueMemberS).Value == reservationID
	})).Return(&dynamodb.DeleteItemOutput{}, nil)

	mockSMTP.On("SendEmailWithGoogle", mock.Anything, mock.Anything, mock.Anything).Return(assert.AnError)

	// Create request body
	requestBody := handlers.RequestChangeRequest{
		Key:    reservationID,
		Code:   "CANCEL",
		Reason: "취소 사유",
	}
	bodyBytes, _ := json.Marshal(requestBody)

	// Create API Gateway request
	ctx := context.Background()
	request := events.APIGatewayV2HTTPRequest{
		Body: string(bodyBytes),
	}

	// Create handler parameters
	params := handlers.RouterHandlerParameters{
		Ctx:        ctx,
		Request:    request,
		DdbClient:  mockDDB,
		SmtpClient: mockSMTP,
	}

	// Call the handler
	response, err := handlers.ManageReservation(params)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, "Failed to send email", response.Body)

	// Verify all mocks were called
	mockDDB.AssertExpectations(t)
	mockSMTP.AssertExpectations(t)
}

func TestManageReservation_ModifyEmailError(t *testing.T) {
	// Setup mock DynamoDB client
	mockDDB := &mocks.MockDDBClient{}
	mockSMTP := &mocks.MockSendEmail{}

	// Mock reservation data
	reservationID := "test-reservation-id"

	// Create time values for the test
	timeValues := []int{13, 14, 15}

	// Mock IsItemExist response with email
	mockDDB.On("GetItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.GetItemInput) bool {
		return *input.TableName == "current_reservation" &&
			input.Key["reservationId"].(*types.AttributeValueMemberS).Value == reservationID
	})).Return(&dynamodb.GetItemOutput{
		Item: map[string]types.AttributeValue{
			"reservationId": &types.AttributeValueMemberS{Value: reservationID},
			"email":         &types.AttributeValueMemberS{Value: "test@example.com"},
			"venueDate":     &types.AttributeValueMemberS{Value: "2025-03-03#studio"},
			"name":          &types.AttributeValueMemberS{Value: "test"},
			"category":      &types.AttributeValueMemberS{Value: "test"},
			"time": &types.AttributeValueMemberL{
				Value: []types.AttributeValue{
					&types.AttributeValueMemberN{Value: "10"},
					&types.AttributeValueMemberN{Value: "11"},
					&types.AttributeValueMemberN{Value: "12"},
				},
			},
		},
	}, nil)

	// Mock UpdateItem response
	mockDDB.On("UpdateItem", mock.Anything, mock.MatchedBy(func(input *dynamodb.UpdateItemInput) bool {
		return *input.TableName == "current_reservation" &&
			input.Key["reservationId"].(*types.AttributeValueMemberS).Value == reservationID
	})).Return(&dynamodb.UpdateItemOutput{}, nil)

	mockSMTP.On("SendEmailWithGoogle", mock.Anything, mock.Anything, mock.Anything).Return(assert.AnError)

	// Create request body
	requestBody := handlers.RequestChangeRequest{
		Key:  reservationID,
		Code: "MODIFY",
		ChangeValues: handlers.ChangeValuesType{
			ChangeTime: timeValues,
			Venue:      "studio",
			Date:       "2025-03-03",
		},
		Reason: "변경 사유",
	}
	bodyBytes, _ := json.Marshal(requestBody)

	// Create API Gateway request
	ctx := context.Background()
	request := events.APIGatewayV2HTTPRequest{
		Body: string(bodyBytes),
	}

	// Create handler parameters
	params := handlers.RouterHandlerParameters{
		Ctx:        ctx,
		Request:    request,
		DdbClient:  mockDDB,
		SmtpClient: mockSMTP,
	}

	// Call the handler
	response, err := handlers.ManageReservation(params)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, "Failed to send email", response.Body)

	// Verify all mocks were called
	mockDDB.AssertExpectations(t)
	mockSMTP.AssertExpectations(t)
}

func TestManageReservation_InvalidJSON(t *testing.T) {
	// Setup mock DynamoDB client
	mockDDB := &mocks.MockDDBClient{}
	mockSMTP := &mocks.MockSendEmail{}

	// Create API Gateway request with invalid JSON
	ctx := context.Background()
	request := events.APIGatewayV2HTTPRequest{
		Body: "{invalid json",
	}

	// Create handler parameters
	params := handlers.RouterHandlerParameters{
		Ctx:        ctx,
		Request:    request,
		DdbClient:  mockDDB,
		SmtpClient: mockSMTP,
	}

	// Call the handler
	response, err := handlers.ManageReservation(params)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, "Failed to parse request body", response.Body)

	// Verify mocks were called
	mockDDB.AssertExpectations(t)
	mockSMTP.AssertExpectations(t)
}
