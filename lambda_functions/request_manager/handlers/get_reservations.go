package handlers

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/sirupsen/logrus"
	"net/http"
	"request_manager/actions"
	"request_manager/response"
)

var log = logrus.New()

type ReservationType struct {
	reservationID string `dynamodbav:"reservationId"`
	category      string `dynamodbav:"category"`
	companion     string `dynamodbav:"companion"`
	email         string `dynamodbav:"email"`
	name          string `dynamodbav:"name"`
	purpose       string `dynamodbav:"purpose"`
	studentID     string `dynamodbav:"studentId"`
	time          []int  `dynamodbav:"time"`
	venueDate     string `dynamodbav:"venueDate"`
}

func GetReservations(params RouterHandlerParameters) (events.APIGatewayV2HTTPResponse, error) {
	ctx := context.Background()
	ddbClient := params.DdbClient

	var scanResults *dynamodb.ScanOutput

	scanResults, err := actions.ScanTable(ctx, ddbClient, "current_reservation")
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
