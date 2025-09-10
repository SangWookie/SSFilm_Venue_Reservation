package handlers

import (
	"context"
	"net/http"
	"request_manager/actions"
	"request_manager/response"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type ReservationType struct {
	ReservationID string `dynamodbav:"reservationId" json:"reservationID"`
	Category      string `dynamodbav:"category" json:"category"`
	Companion     string `dynamodbav:"companion" json:"companion"`
	Email         string `dynamodbav:"email" json:"email"`
	Name          string `dynamodbav:"name" json:"name"`
	Purpose       string `dynamodbav:"purpose" json:"purpose"`
	StudentID     string `dynamodbav:"studentId" json:"studentID"`
	Time          []int  `dynamodbav:"time" json:"time"`
	VenueDate     string `dynamodbav:"venueDate" json:"venueDate"`
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

func GetReservationV2(params RouterHandlerParameters) (events.APIGatewayV2HTTPResponse, error) {
	yearMonth := params.Request.QueryStringParameters["year_month"]

	keyCond := expression.KeyAnd(
		expression.KeyEqual(expression.Key("date"), expression.Value(yearMonth)),
		expression.KeyBeginsWith(expression.Key("venueDate"), yearMonth),
	)
	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()
	if err != nil {
		return response.APIGatewayResponseError("Build expression error", 500), err
	}

	out, err := params.DdbClient.Query(params.Ctx, &dynamodb.QueryInput{
		TableName:                 aws.String("current_reservation"),
		IndexName:                 aws.String("date-venueDate-index"),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeValues: expr.Values(),
	})
	if err != nil {
		return response.APIGatewayResponseError("DDB Query Error", 500), err
	}

	var reservations []ReservationType
	err = attributevalue.UnmarshalListOfMaps(out.Items, &reservations)
	if err != nil {
		return response.APIGatewayResponseError("Unmarshal Error", 500), err
	}

	return response.APIGatewayResponseOK(reservations, http.StatusOK), nil
}
