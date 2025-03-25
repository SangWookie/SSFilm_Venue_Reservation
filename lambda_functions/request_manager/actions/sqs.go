package actions

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"os"
)

var queueURL = os.Getenv("QUEUE_URL")

type SQSClientiface interface {
	SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
}

type SQSClient struct {
	Client *sqs.Client
}

func (r *SQSClient) SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error) {
	return r.Client.SendMessage(ctx, params, optFns...)
}

func SendEmail(ctx context.Context, sqsClient SQSClientiface, email, emailType string, data ReservationEmailData) error {
	message := map[string]interface{}{
		"data":  data,
		"type":  emailType,
		"email": email,
	}

	body, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("failed to marshal message body, %v", err)
	}

	_, err = sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(string(body)),
	})

	if err != nil {
		log.Fatalf("failed to send message, %v", err)
	}

	return nil
}
