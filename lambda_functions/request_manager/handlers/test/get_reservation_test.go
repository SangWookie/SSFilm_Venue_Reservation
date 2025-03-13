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

func TestGetReservations_Success(t *testing.T) {
	mockDDB := &mocks.MockDDBClient{}
	mockDDB.On("Scan", mock.Anything, mock.Anything).Return(&dynamodb.ScanOutput{
		Items: []map[string]types.AttributeValue{
			{
				"reservationId": &types.AttributeValueMemberS{Value: "0c6cb985d405b..."},
				"category":      &types.AttributeValueMemberS{Value: "수업"},
				"companion":     &types.AttributeValueMemberS{Value: "james, andrew"},
				"email":         &types.AttributeValueMemberS{Value: "tester@tester.com"},
				"name":          &types.AttributeValueMemberS{Value: "tester"},
				"purpose":       &types.AttributeValueMemberS{Value: "specific purpose for usage"},
				"studentId":     &types.AttributeValueMemberN{Value: "20201728"},
				"time": &types.AttributeValueMemberL{
					Value: []types.AttributeValue{
						&types.AttributeValueMemberN{Value: "10"},
						&types.AttributeValueMemberN{Value: "11"},
						&types.AttributeValueMemberN{Value: "12"},
					},
				},
				"venueDate": &types.AttributeValueMemberS{Value: "2025-03-31#studio"},
			},
		},
	}, nil)

	ctx := context.Background()
	request := events.APIGatewayV2HTTPRequest{}

	response, err := handlers.GetReservations(ctx, request, mockDDB)

	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)
}
