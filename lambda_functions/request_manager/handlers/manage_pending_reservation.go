package handlers

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"net/http"
	"request_manager/actions"
	"request_manager/response"
	"strings"
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
	smtpClient := params.SmtpClient

	var reqBody RequestDeleteType
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		log.Errorln(err)
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

				// TODO 시간 변경을 어떻게 보여주지?
				emailContent, err := actions.GetReservationCompleteTemplate(actions.ReservationEmailData{
					Name:     isExist["name"].(*types.AttributeValueMemberS).Value,
					Location: room,
					Time:     date,
					Category: isExist["category"].(*types.AttributeValueMemberS).Value,
					Details:  reqBody.Reason,
				})
				err = actions.SendEmail(smtpClient, emailValue.Value, "예약 확인", emailContent)
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

				// TODO 시간 변경을 어떻게 보여주지?
				emailContent, err := actions.GetReservationCanceledTemplate(actions.ReservationEmailData{
					Name:     isExist["name"].(*types.AttributeValueMemberS).Value,
					Location: room,
					Time:     date,
					Category: isExist["category"].(*types.AttributeValueMemberS).Value,
					Details:  "",
				})
				err = actions.SendEmail(smtpClient, emailValue.Value, "관리자 예약 반려", emailContent)
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
