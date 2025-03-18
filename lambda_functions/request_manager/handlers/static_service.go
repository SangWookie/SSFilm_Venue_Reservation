package handlers

import (
	"context"
	"encoding/json"
	"fmt"
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

	if reqBody.Venue != "" {
		// TODO 방 통계
		fmt.Printf("Venue: %s\n", reqBody.Venue)

		result, err := actions.GetReservationsWithVenue(ctx, ddbClient, "current_reservation", reqBody.Venue, reqBody.Month)

		if err != nil {
			return response.APIGatewayResponseError(err.Error(), http.StatusBadRequest), nil
		}

		return response.APIGatewayResponseOK(result, http.StatusOK), nil
	} else {
		// TODO 학생 통계
		fmt.Printf("StudentID: %s\n", reqBody.StudentID)

		result, err := actions.GetReservationsWithStudentID(ctx, ddbClient, "current_reservation", reqBody.StudentID, reqBody.Month)

		if err != nil {
			return response.APIGatewayResponseError(err.Error(), http.StatusBadRequest), nil
		}

		return response.APIGatewayResponseOK(result, http.StatusOK), nil
	}
}
