package handlers

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"net/http"
	"request_manager/actions"
	"request_manager/response"

	"github.com/aws/aws-lambda-go/events"
)

type StaticRequest struct {
	Month string `json:"month"`
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

	executionResults, err = actions.GetHistory(ctx, ddbClient, "current_reservation", reqBody.Month)

	if err != nil {
		return response.APIGatewayResponseError(err.Error(), http.StatusBadRequest), nil
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
