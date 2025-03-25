package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
)

var (
	log        = logrus.New()
	smtpClient = &SMTPManager{}
)

type STMPRequestBodyType struct {
	Type  string               `json:"type"`
	Data  ReservationEmailData `json:"data"`
	Email string               `json:"email"`
}

func processMessage(record events.SQSMessage) error {
	log.Info("Processed message: %s\n", record.Body)

	var reqBody STMPRequestBodyType
	err := json.Unmarshal([]byte(record.Body), &reqBody)
	if err != nil {
		return err
	}

	var emailContent string
	var subject string
	switch reqBody.Type {
	case "ACCEPT":
		emailContent, err = GetReservationCompleteTemplate(reqBody.Data)
		subject = "예약 확정 안내"
	case "MODIFY":
		emailContent, err = GetReservationModifiedTemplate(reqBody.Data)
		subject = "관리자 예약 수정 안내"
	case "CANCEL":
		emailContent, err = GetReservationCanceledTemplate(reqBody.Data)
		subject = "관리자 예약 취소 안내"
	case "DENY":
		emailContent, err = GetRequestDenyTemplate(reqBody.Data)
		subject = "관리자 예약 거부 안내"
	default:
		return err
	}

	smtpErr := SendEmail(smtpClient, reqBody.Email, subject, emailContent)
	if smtpErr != nil {
		return smtpErr
	}

	return nil
}
func handleRequest(event events.SQSEvent) error {
	for _, record := range event.Records {
		err := processMessage(record)
		if err != nil {
			return err
		}
	}
	fmt.Println("done")
	return nil
}

func main() {
	lambda.Start(handleRequest)
}
