package handlers

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"net/http"
	"request_manager/actions"
	"request_manager/response"
)

type RequestDeleteType struct {
	Key  string `json:"requestID"`
	Code string `json:"code"`
}

func ManagePendingReservation(ctx context.Context, request events.APIGatewayV2HTTPRequest, ddbClient actions.DDBClientiface) (events.APIGatewayV2HTTPResponse, error) {
	// todo 예약 승인 대기 중인 테이블에서 승인 또는 거부
	/**

	1. 처리 대상 확인
	2. 처리 목표 확인
	3. 처리
	3. pending table에서 삭제
	*/

	var reqBody RequestDeleteType
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
	}

	key := map[string]types.AttributeValue{
		"requestId": &types.AttributeValueMemberS{Value: reqBody.Key},
	}

	isExist, err := actions.IsItemExist(ctx, ddbClient, "pending_reservation", key)
	if err != nil || !isExist {
		return response.APIGatewayResponseError("Not found Item", http.StatusNotFound), nil
	}

	pendedReservation, err := actions.GetPendingItem(ctx, ddbClient, key)

	switch reqBody.Code {
	case "ACCEPT":
		// current_reservation 에 예약 정보추가
		err := actions.AcceptReservation(ctx, ddbClient, pendedReservation)
		if err != nil {
		}
	case "DENY":
		// 사용자에게 취소 이메일 발송
	default:
		return response.APIGatewayResponseError("Not Found", http.StatusNotFound), nil
	}

	err = actions.DeletePendingItem(ctx, ddbClient, key)
	if err != nil {
	}

	return response.APIGatewayResponseOK("success", http.StatusOK), nil
}
