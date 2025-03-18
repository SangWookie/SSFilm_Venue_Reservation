package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"request_manager/actions"
	"request_manager/response"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type StaticRequest struct {
	StudentID string `json:"student_id,omitempty"`
	Venue     string `json:"venue,omitempty"`
	Month     string `json:"month"`
}

func GetStatic(params RouterHandlerParameters) (events.APIGatewayV2HTTPResponse, error) {
	request := params.Request
	ctx := context.Background()
	ddbClient := params.DdbClient

	var reqBody StaticRequest
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		return response.APIGatewayResponseError("Failed to parse request body", http.StatusBadRequest), nil
	}

	if reqBody.Venue != "" {
		// TODO 방 통계
		fmt.Printf("Venue: %s\n", reqBody.Venue)

		result, err := actions.GetQueryResult(ctx, ddbClient, &dynamodb.QueryInput{
			TableName:              aws.String("current_reservation"),
			KeyConditionExpression: aws.String("begins_with(venueDate, :date)"),
			FilterExpression:       aws.String("contains(venueDate, :room)"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":date": &types.AttributeValueMemberS{Value: reqBody.Month},
				":room": &types.AttributeValueMemberS{Value: reqBody.Venue},
			},
		})

		if err != nil {
			return response.APIGatewayResponseError(err.Error(), http.StatusBadRequest), nil
		}

		return response.APIGatewayResponseOK(result, http.StatusOK), nil
	} else {
		// TODO 학생 통계
		fmt.Printf("StudentID: %s\n", reqBody.StudentID)

		result, err := actions.GetQueryResult(ctx, ddbClient, &dynamodb.QueryInput{
			TableName:              aws.String("current_reservation"),
			KeyConditionExpression: aws.String("venueDate begins_with :date and studentId = :studentId"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":date":      &types.AttributeValueMemberS{Value: reqBody.Month},
				":studentId": &types.AttributeValueMemberS{Value: reqBody.StudentID},
			},
		})

		if err != nil {
			return response.APIGatewayResponseError(err.Error(), http.StatusBadRequest), nil
		}

		return response.APIGatewayResponseOK(result, http.StatusOK), nil
	}
}
