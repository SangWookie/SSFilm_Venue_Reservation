package handlers

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
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

	var reservations []ReservationType
	for _, scanResult := range scanResults.Items {
		var tmp ReservationType
		err = attributevalue.UnmarshalMap(scanResult, &tmp)
		if err != nil {
			log.Errorln("Parser Error", err)
			return response.APIGatewayResponseError("Internal Server Error", 500), err
		}

		reservations = append(reservations, tmp)
	}
	return response.APIGatewayResponseOK(reservations, http.StatusOK), nil
}
