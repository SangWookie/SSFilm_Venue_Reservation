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

func TestGetStaticVenue(t *testing.T) {
	// Setup mock DynamoDB client
	mockDDB := new(mocks.MockDDBClient)
	mockSMTP := new(mocks.MockSendEmail)

	// Mock Query response for venue statistics
	mockDDB.On("Query", mock.Anything, mock.MatchedBy(func(input *dynamodb.QueryInput) bool {
		return *input.TableName == "current_reservation"
	})).Return(&dynamodb.QueryOutput{
		Items: []map[string]types.AttributeValue{
			{
				"venueDate": &types.AttributeValueMemberS{Value: "2024-03-Room1"},
				"venue":     &types.AttributeValueMemberS{Value: "Room1"},
			},
		},
	}, nil)

	// Create request body
	requestBody := handlers.StaticRequest{
		Venue: "Room1",
		Month: "2024-03",
	}
	bodyBytes, _ := json.Marshal(requestBody)

	// Create API Gateway request
	request := events.APIGatewayV2HTTPRequest{
		Body: string(bodyBytes),
	}

	// Create handler parameters
	params := handlers.RouterHandlerParameters{
		Ctx:        context.Background(),
		Request:    request,
		DdbClient:  mockDDB,
		SmtpClient: mockSMTP,
	}

	// Call the handler
	response, err := handlers.GetStatic(params)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)

	// Verify mock was called
	mockDDB.AssertExpectations(t)
}

func TestGetStaticStudent(t *testing.T) {
	// Setup mock DynamoDB client
	mockDDB := new(mocks.MockDDBClient)
	mockSMTP := new(mocks.MockSendEmail)

	// Mock Query response for student statistics
	mockDDB.On("Query", mock.Anything, mock.MatchedBy(func(input *dynamodb.QueryInput) bool {
		return *input.TableName == "current_reservation"
	})).Return(&dynamodb.QueryOutput{
		Items: []map[string]types.AttributeValue{
			{
				"venueDate": &types.AttributeValueMemberS{Value: "2024-03-Room1"},
				"studentId": &types.AttributeValueMemberS{Value: "12345"},
			},
		},
	}, nil)

	// Create request body
	requestBody := handlers.StaticRequest{
		StudentID: "12345",
		Month:     "2024-03",
	}
	bodyBytes, _ := json.Marshal(requestBody)

	// Create API Gateway request
	request := events.APIGatewayV2HTTPRequest{
		Body: string(bodyBytes),
	}

	// Create handler parameters
	params := handlers.RouterHandlerParameters{
		Ctx:        context.Background(),
		Request:    request,
		DdbClient:  mockDDB,
		SmtpClient: mockSMTP,
	}

	// Call the handler
	response, err := handlers.GetStatic(params)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)

	// Verify mock was called
	mockDDB.AssertExpectations(t)
}
