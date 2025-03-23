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
type RequestChangeTimeRequest struct {
	Key          string           `json:"reservationID"`
	Code         string           `json:"code"`
	ChangeValues ChangeValuesType `json:"changeValues"`
	Reason       string           `json:"reason"`
}

func ManageReservation(params RouterHandlerParameters) (events.APIGatewayV2HTTPResponse, error) {
	request := params.Request
	ctx := context.Background()
	ddbClient := params.DdbClient
	smtpClient := params.SmtpClient

	var reqBody RequestChangeTimeRequest
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

				// TODO 시간 변경을 어떻게 보여주지?
				emailContent, err := actions.GetReservationCanceledTemplate(actions.ReservationEmailData{
					Name:     reservationItem["name"].(*types.AttributeValueMemberS).Value,
					Location: room,
					Time:     date,
					Category: reservationItem["category"].(*types.AttributeValueMemberS).Value,
					Details:  reqBody.Reason,
				})
				err = actions.SendEmail(smtpClient, emailValue.Value, "예약 취소 확인", emailContent)
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
		// 수동으로 ChangeTime 값을 숫자 리스트로 변환
		var changeTimeList []types.AttributeValue
		for _, t := range reqBody.ChangeValues.ChangeTime {
			changeTimeList = append(changeTimeList, &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", t)})
		}
		changeValuesMap["ChangeTime"] = &types.AttributeValueMemberL{Value: changeTimeList}

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
				// TODO 시간 변경을 어떻게 보여주지?
				emailContent, err := actions.GetReservationModifiedTemplate(actions.ReservationEmailData{
					Name:     reservationItem["name"].(*types.AttributeValueMemberS).Value,
					Location: room,
					Time:     time,
					Category: reservationItem["category"].(*types.AttributeValueMemberS).Value,
					Details:  reqBody.Reason,
				})
				err = actions.SendEmail(smtpClient, emailValue.Value, "관리자 예약 시간 변경", emailContent)
				if err != nil {
					return response.APIGatewayResponseError("Failed to send email", http.StatusInternalServerError), nil
				}
			}
		}
	}

	return response.APIGatewayResponseOK("success", http.StatusOK), nil
}
