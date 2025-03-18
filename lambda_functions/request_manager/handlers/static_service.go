package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"net/http"
	"request_manager/actions"
	"request_manager/response"

	"github.com/aws/aws-lambda-go/events"
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

	var executionResults *dynamodb.ExecuteStatementOutput
	var result []ReservationType

	if reqBody.Venue != "" {
		// TODO 방 통계
		fmt.Printf("Venue: %s\n", reqBody.Venue)

		executionResults, err = actions.GetReservationsWithVenue(ctx, ddbClient, "current_reservation", reqBody.Venue, reqBody.Month)

		if err != nil {
			return response.APIGatewayResponseError(err.Error(), http.StatusBadRequest), nil
		}

	} else {
		// TODO 학생 통계
		fmt.Printf("StudentID: %s\n", reqBody.StudentID)

		executionResults, err = actions.GetReservationsWithStudentID(ctx, ddbClient, "current_reservation", reqBody.StudentID, reqBody.Month)

		if err != nil {
			return response.APIGatewayResponseError(err.Error(), http.StatusBadRequest), nil
		}

	}

	for _, executionResult := range executionResults.Items {
		var tmp ReservationType
		err = attributevalue.UnmarshalMap(executionResult, &tmp)
		if err != nil {
			log.Errorln("Parser Error", err)
			return response.APIGatewayResponseError("Internal Server Error", 500), err
		}

		result = append(result, tmp)
	}

	return response.APIGatewayResponseOK(result, http.StatusOK), nil
}
