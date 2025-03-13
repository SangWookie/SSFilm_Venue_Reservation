package handlers

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"net/http"
	"request_manager/actions"
	"request_manager/response"
)

func GetReservations(ctx context.Context, request events.APIGatewayV2HTTPRequest, ddbClient actions.DDBClientiface) (events.APIGatewayV2HTTPResponse, error) {
	var scanResults []map[string]types.AttributeValue

	scanResults, err := actions.ScanTable(ctx, ddbClient, "current_reservation")
	if err != nil {
		return response.APIGatewayResponseError("Internal Server Error", 500), err
	}

	return response.APIGatewayResponseOK(scanResults, http.StatusOK), nil
}
