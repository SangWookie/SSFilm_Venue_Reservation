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

type VenueInfoTYpe struct {
	Venue       string `dynamodbav:"venue" json:"venue"`
	AllowPolicy string `dynamodbav:"allowPolicy" json:"allowPolicy"`
	VenueKor    string `dynamodbav:"venueKor" json:"venueKor"`
}

func GetVenueInfo(params RouterHandlerParameters) (events.APIGatewayV2HTTPResponse, error) {
	ctx := context.Background()
	ddbClient := params.DdbClient

	var scanResults *dynamodb.ScanOutput

	scanResults, err := actions.ScanTable(ctx, ddbClient, "venue_info")
	if err != nil {
		return response.APIGatewayResponseError("Internal Server Error", 500), err
	}

	var reservations []VenueInfoTYpe
	for _, scanResult := range scanResults.Items {
		var tmp VenueInfoTYpe
		err = attributevalue.UnmarshalMap(scanResult, &tmp)
		if err != nil {
			log.Errorln("Parser Error", err)
			return response.APIGatewayResponseError("Internal Server Error", 500), err
		}

		reservations = append(reservations, tmp)
	}

	return response.APIGatewayResponseOK(reservations, http.StatusOK), nil
}
