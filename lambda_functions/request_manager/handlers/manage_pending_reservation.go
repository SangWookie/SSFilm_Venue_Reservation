package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"request_manager/actions"
	"request_manager/response"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type RequestDeleteType struct {
	Key    string `json:"requestID"`
	Code   string `json:"code"`
	Reason string `json:"reason"`
}

func ManagePendingReservation(params RouterHandlerParameters) (events.APIGatewayV2HTTPResponse, error) {
	request := params.Request
	ctx := context.Background()
	ddbClient := params.DdbClient
	sqsClient := params.SQSClient

	var reqBody RequestDeleteType
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		return response.APIGatewayResponseError("Failed to parse request body", http.StatusBadRequest), nil
	}

	key := map[string]types.AttributeValue{
		"requestId": &types.AttributeValueMemberS{Value: reqBody.Key},
	}

	isExist, err := actions.IsItemExist(ctx, ddbClient, "pending_reservation", key)
	if err != nil {
		return response.APIGatewayResponseError("Not found Item", http.StatusNotFound), nil
	}

	pendedReservation, err := actions.GetPendingItem(ctx, ddbClient, key)

	switch reqBody.Code {
	case "ACCEPT":
		// current_reservation 에 예약 정보추가
		err := actions.AcceptReservation(ctx, ddbClient, pendedReservation)
		if err != nil {
			return response.APIGatewayResponseError("Not found Item", http.StatusNotFound), nil
		}

		// 이메일 추출 및 전송
		if emailAttr, ok := isExist["email"]; ok {
			if emailValue, ok := emailAttr.(*types.AttributeValueMemberS); ok {
				venueDate := isExist["venueDate"].(*types.AttributeValueMemberS).Value
				date := strings.Split(venueDate, "#")[0]
				room := strings.Split(venueDate, "#")[1]

				// Get time from time list
				timeList := isExist["time"].(*types.AttributeValueMemberL).Value
				startTime := timeList[0].(*types.AttributeValueMemberN).Value
				endTime := timeList[len(timeList)-1].(*types.AttributeValueMemberN).Value
				time := fmt.Sprintf("%s [%s - %s]", date, startTime, endTime)

				// Prepare email data
				emailData := actions.ReservationEmailData{
					Name:     isExist["name"].(*types.AttributeValueMemberS).Value,
					Location: room,
					Time:     time,
					Category: isExist["category"].(*types.AttributeValueMemberS).Value,
					Details:  reqBody.Reason,
				}

				// Send email via SQS
				err = actions.SendEmail(ctx, sqsClient, emailValue.Value, reqBody.Code, emailData)
				if err != nil {
					return response.APIGatewayResponseError("Failed to send email", http.StatusInternalServerError), nil
				}
			}
		}
	case "DENY":
		// 이메일 추출 및 전송
		if emailAttr, ok := isExist["email"]; ok {
			if emailValue, ok := emailAttr.(*types.AttributeValueMemberS); ok {
				venueDate := isExist["venueDate"].(*types.AttributeValueMemberS).Value
				date := strings.Split(venueDate, "#")[0]
				room := strings.Split(venueDate, "#")[1]

				// Get time from time list
				timeList := isExist["time"].(*types.AttributeValueMemberL).Value
				startTime := timeList[0].(*types.AttributeValueMemberN).Value
				endTime := timeList[len(timeList)-1].(*types.AttributeValueMemberN).Value
				time := fmt.Sprintf("%s [%s - %s]", date, startTime, endTime)

				// Prepare email data
				emailData := actions.ReservationEmailData{
					Name:     isExist["name"].(*types.AttributeValueMemberS).Value,
					Location: room,
					Time:     time,
					Category: isExist["category"].(*types.AttributeValueMemberS).Value,
					Details:  reqBody.Reason,
				}

				// Send email via SQS
				err = actions.SendEmail(ctx, sqsClient, emailValue.Value, reqBody.Code, emailData)
				if err != nil {
					return response.APIGatewayResponseError("Failed to send email", http.StatusInternalServerError), nil
				}
			}
		}
	default:
		return response.APIGatewayResponseError("Not Found", http.StatusNotFound), nil
	}

	err = actions.DeletePendingItem(ctx, ddbClient, key)
	if err != nil {
		return response.APIGatewayResponseError("Failed to delete pending item", http.StatusInternalServerError), nil
	}

	return response.APIGatewayResponseOK("success", http.StatusOK), nil
}
