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
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ChangeValuesType struct {
	ChangeTime []int  `json:"changeTime"`
	Venue      string `json:"venue"`
	Date       string `json:"date"`
}

type RequestChangeRequest struct {
	Key          string           `json:"reservationID"`
	Code         string           `json:"code"`
	ChangeValues ChangeValuesType `json:"changeValues"`
	Reason       string           `json:"reason"`
}

func ManageReservation(params RouterHandlerParameters) (events.APIGatewayV2HTTPResponse, error) {
	request := params.Request
	ctx := context.Background()
	ddbClient := params.DdbClient
	sqsClient := params.SQSClient

	var reqBody RequestChangeRequest
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		return response.APIGatewayResponseError("Failed to parse request body", http.StatusBadRequest), nil
	}

	key := map[string]types.AttributeValue{
		"reservationId": &types.AttributeValueMemberS{Value: reqBody.Key},
	}

	reservationItem, err := actions.IsItemExist(ctx, ddbClient, "current_reservation", key)
	if err != nil {
		return response.APIGatewayResponseError("Not found Item", http.StatusNotFound), nil
	}

	switch reqBody.Code {
	case "CANCEL":
		// 예약 취소
		err := actions.DeleteReservationItem(ctx, ddbClient, key)
		if err != nil {
			return response.APIGatewayResponseError("Failed to cancel reservation", http.StatusInternalServerError), nil
		}

		// 이메일 추출 및 전송
		if emailAttr, ok := reservationItem["email"]; ok {
			if emailValue, ok := emailAttr.(*types.AttributeValueMemberS); ok {
				venueDate := reservationItem["venueDate"].(*types.AttributeValueMemberS).Value
				date := strings.Split(venueDate, "#")[0]
				room := strings.Split(venueDate, "#")[1]

				// SQS에 이메일 요청 전송
				emailReq := actions.ReservationEmailData{
					Name:     reservationItem["name"].(*types.AttributeValueMemberS).Value,
					Location: room,
					Time:     date,
					Category: reservationItem["category"].(*types.AttributeValueMemberS).Value,
					Details:  reqBody.Reason,
				}

				err = actions.SendEmail(ctx, sqsClient, emailValue.Value, reqBody.Code, emailReq)
				if err != nil {
					return response.APIGatewayResponseError("Failed to send email", http.StatusInternalServerError), nil
				}
			}
		}
	case "MODIFY":
		// 예약 시간 변경
		changeValuesMap, err := attributevalue.MarshalMap(reqBody.ChangeValues)
		if err != nil {
			return response.APIGatewayResponseError("Failed to marshal change values", http.StatusInternalServerError), nil
		}
		err = actions.ChangeReservationValues(ctx, ddbClient, key, changeValuesMap)
		if err != nil {
			return response.APIGatewayResponseError("Failed to modify reservation time", http.StatusInternalServerError), nil
		}

		// 이메일 추출 및 전송
		if emailAttr, ok := reservationItem["email"]; ok {
			if emailValue, ok := emailAttr.(*types.AttributeValueMemberS); ok {
				date := reqBody.ChangeValues.Date
				room := reqBody.ChangeValues.Venue
				time := fmt.Sprintf("%s [%d - %d]", date, reqBody.ChangeValues.ChangeTime[0], reqBody.ChangeValues.ChangeTime[len(reqBody.ChangeValues.ChangeTime)-1])

				// SQS에 이메일 요청 전송
				emailReq := actions.ReservationEmailData{
					Name:     reservationItem["name"].(*types.AttributeValueMemberS).Value,
					Location: room,
					Time:     time,
					Category: reservationItem["category"].(*types.AttributeValueMemberS).Value,
					Details:  reqBody.Reason,
				}

				err = actions.SendEmail(ctx, sqsClient, emailValue.Value, reqBody.Code, emailReq)
				if err != nil {
					return response.APIGatewayResponseError("Failed to send email", http.StatusInternalServerError), nil
				}
			}
		}
	}

	return response.APIGatewayResponseOK("success", http.StatusOK), nil
}
