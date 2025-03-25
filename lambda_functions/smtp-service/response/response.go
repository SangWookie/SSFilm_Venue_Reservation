package response

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var (
	header = map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET, POST, OPTIONS",
		"Access-Control-Allow-Headers": "Content-Type",
	}
)

func APIGatewayResponseError(msg string, code int) events.APIGatewayV2HTTPResponse {
	log.Errorf(msg)
	return events.APIGatewayV2HTTPResponse{
		StatusCode: code,
		Body:       msg,
	}
}

func APIGatewayResponseOK(body any, code int) events.APIGatewayV2HTTPResponse {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return APIGatewayResponseError("Internal Server Error", 500)
	}
	return events.APIGatewayV2HTTPResponse{
		StatusCode: code,
		Headers:    header,
		Body:       string(jsonBody),
	}
}
