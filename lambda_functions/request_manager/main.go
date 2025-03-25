package main

import (
	"context"
	"request_manager/actions"
	"request_manager/handlers"
	"request_manager/response"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

var (
	sdkConfig, configErr = config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("ap-northeast-2"),
	)

	// 람다 최적화를 위한 전역 변수 설정
	ddbClient = &actions.DDBClient{
		DynamoDbClient: dynamodb.NewFromConfig(sdkConfig),
	}
	sqsClient = &actions.SQSClient{
		Client: sqs.NewFromConfig(sdkConfig),
	}
)

type RouteHandler func(params handlers.RouterHandlerParameters) (events.APIGatewayV2HTTPResponse, error)

var routes = map[string]RouteHandler{
	"getPendingReservations":   handlers.GetPendingReservations,
	"getReservations":          handlers.GetReservations,
	"manageReservation":        handlers.ManageReservation,
	"managePendingReservation": handlers.ManagePendingReservation,
	"statistic":                handlers.GetStatic,
	"getVenueInfo":             handlers.GetVenueInfo,
}

func handleRequest(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	if configErr != nil {
		return response.APIGatewayResponseError("Not found configuration values", 400), nil
	}
	log.Info("New request received1")

	var path = request.PathParameters["proxy"]

	log.Info("Path parameters: ", path)
	log.Info("Client IP: ", request.RequestContext.HTTP.SourceIP)

	routerHandlerParameters := handlers.RouterHandlerParameters{
		Ctx:       ctx,
		Request:   request,
		DdbClient: ddbClient,
		SQSClient: sqsClient,
	}

	if handler, exist := routes[path]; exist {
		return handler(routerHandlerParameters)
	}

	return response.APIGatewayResponseError("Internal Server Error", 500), nil
}

func main() {
	lambda.Start(handleRequest)
}
