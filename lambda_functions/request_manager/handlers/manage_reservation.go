package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"request_manager/actions"
	"request_manager/response"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// TimeValue represents the JSON structure for time values
type TimeValue struct {
	N string `json:"N"`
}

type RequestChangeTimeRequest struct {
	Key        string      `json:"reservationID"`
	Code       string      `json:"code"`
	ChangeTime []TimeValue `json:"changeTime"`
}

func ManageReservation(ctx context.Context, request events.APIGatewayV2HTTPRequest, ddbClient actions.DDBClientiface) (events.APIGatewayV2HTTPResponse, error) {
	var reqBody RequestChangeTimeRequest
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		return response.APIGatewayResponseError("Failed to parse request body", http.StatusBadRequest), nil
	}

	key := map[string]types.AttributeValue{
		"reservationId": &types.AttributeValueMemberS{Value: reqBody.Key},
	}

	// Convert the TimeValue slice to AttributeValue slice
	requestChangeTime := make([]types.AttributeValue, len(reqBody.ChangeTime))
	for i, tv := range reqBody.ChangeTime {
		requestChangeTime[i] = &types.AttributeValueMemberN{Value: tv.N}
	}

	isExist, err := actions.IsItemExist(ctx, ddbClient, "current_reservation", key)
	if err != nil || !isExist {
		return response.APIGatewayResponseError("Not found Item", http.StatusNotFound), nil
	}

	switch reqBody.Code {
	case "CANCEL":
		// 예약 취소
		err := actions.DeleteReservationItem(ctx, ddbClient, key)
		if err != nil {
			return response.APIGatewayResponseError("Failed to cancel reservation", http.StatusInternalServerError), nil
		}
	case "MODIFY":
		// 예약 시간 변경
		err := actions.ChangeReservationTime(ctx, ddbClient, key, requestChangeTime)
		if err != nil {
			return response.APIGatewayResponseError("Failed to modify reservation time", http.StatusInternalServerError), nil
		}
	}

	return response.APIGatewayResponseOK("success", http.StatusOK), nil
}
