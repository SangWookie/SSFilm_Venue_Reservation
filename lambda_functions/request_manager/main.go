package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
	"request_manager/response"
)

var log = logrus.New()

type RouteHandler func(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error)

var routes = map[string]RouteHandler{}

func handleRequest(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	log.Info("New request received1")

	var path = request.PathParameters["proxy"]

	log.Info("Path parameters: ", path)
	log.Info("Client IP: ", request.RequestContext.HTTP.SourceIP)

	if handler, exist := routes[path]; exist {
		return handler(request)
	}

	return response.APIGatewayResponseError("Internal Server Error", 500), nil
}

func main() {
	lambda.Start(handleRequest)
}
