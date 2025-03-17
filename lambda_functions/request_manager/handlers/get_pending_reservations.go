package handlers

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"net/http"
	"request_manager/actions"
	"request_manager/response"
)

func GetPendingReservations(params RouterHandlerParameters) (events.APIGatewayV2HTTPResponse, error) {
	ctx := context.Background()
	ddbClient := params.DdbClient

	var scanResults *dynamodb.ScanOutput
	scanResults, err := actions.ScanTable(ctx, ddbClient, "pending_reservation")
	if err != nil {
		return response.APIGatewayResponseError("Internal Server Error", 500), err
	}

	return response.APIGatewayResponseOK(scanResults, http.StatusOK), nil
}
