package mode_manager

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sirupsen/logrus"
	"main/response"
	"net/http"
)

var log = logrus.New()

var (
	sdkConfig, configErr = config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("ap-northeast-2"),
	)

	// 람다 최적화를 위한 전역 변수 설정
	ddbClient = dynamodb.NewFromConfig(sdkConfig)

	tableName = "venue_info"

	ctx = context.Background()
)

type RequestChangeMode struct {
	Key  map[string]types.AttributeValue `json:"key"`
	Mode string                          `json:"mode"`
}

func changeMod(key map[string]types.AttributeValue, mode string) (events.APIGatewayV2HTTPResponse, error) {
	update := expression.Set(expression.Name("allowPolicy"), expression.Value(mode))

	// 업데이트 표현식 생성
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		log.Printf("Couldn't build expression for update. Error: %v\n", err)
		return response.APIGatewayResponseError("Internal Server Error", 500), err
	}

	_, err = ddbClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(tableName),
		Key:                       key,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              types.ReturnValueUpdatedNew,
	})

	if err != nil {
		return response.APIGatewayResponseError("Internal Server Error", 500), err
	}

	return response.APIGatewayResponseOK("success", http.StatusOK), nil
}
func handleRequest(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	if configErr != nil {
		return response.APIGatewayResponseError("Not found configuration values", 400), nil
	}

	var reqBody RequestChangeMode
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
	}

	return changeMod(reqBody.Key, reqBody.Mode)
}
func main() {
	lambda.Start(handleRequest)
}
