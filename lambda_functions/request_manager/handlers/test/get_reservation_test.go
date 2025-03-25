package test

import (
	"context"
	"request_manager/handlers"
	"request_manager/mocks"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetReservations(t *testing.T) {
	// Setup mock DynamoDB client
	mockDDB := &mocks.MockDDBClient{}
	mockSQS := &mocks.MockSQSClient{}

	// Mock Scan response
	mockDDB.On("Scan", mock.Anything, mock.MatchedBy(func(input *dynamodb.ScanInput) bool {
		return *input.TableName == "current_reservation"
	})).Return(&dynamodb.ScanOutput{
		Items: []map[string]types.AttributeValue{
			{
				"reservationId": &types.AttributeValueMemberS{Value: "test-id"},
				"email":         &types.AttributeValueMemberS{Value: "test@example.com"},
			},
		},
	}, nil)

	// Create API Gateway request
	ctx := context.Background()
	request := events.APIGatewayV2HTTPRequest{}

	// Create handler parameters
	params := handlers.RouterHandlerParameters{
		Ctx:       ctx,
		Request:   request,
		DdbClient: mockDDB,
		SQSClient: mockSQS,
	}

	// Call the handler
	response, err := handlers.GetReservations(params)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)

	// Verify mock was called
	mockDDB.AssertExpectations(t)
	mockSQS.AssertExpectations(t)
}
