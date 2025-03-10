package main

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"request_manager/response"
)

func FetchPendingReservations(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	tmp := map[string]interface{}{
		"msg": "This is fetching pending user list test",
	}
	return response.APIGatewayResponseOK(tmp, http.StatusOK), nil
}
func HandlePendingReservations(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	tmp := map[string]interface{}{
		"msg": "This is fetching pending user list test ",
	}
	return response.APIGatewayResponseOK(tmp, http.StatusOK), nil
}
func ManageReservation(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	tmp := map[string]interface{}{
		"msg": "This is managing user list tset",
	}
	return response.APIGatewayResponseOK(tmp, http.StatusOK), nil
}
func ModifyReservation(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	tmp := map[string]interface{}{
		"msg": "This is modifying user list test",
	}
	return response.APIGatewayResponseOK(tmp, http.StatusOK), nil
}
