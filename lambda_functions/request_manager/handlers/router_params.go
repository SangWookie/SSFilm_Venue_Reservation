package handlers

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"request_manager/actions"
)

type RouterHandlerParameters struct {
	Ctx        context.Context
	Request    events.APIGatewayV2HTTPRequest
	DdbClient  actions.DDBClientiface
	SmtpClient actions.SMTPManagerIFace
}
