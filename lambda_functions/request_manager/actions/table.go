package actions

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type DDBClientiface interface {
	Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)
	UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
	ExecuteStatement(ctx context.Context, params *dynamodb.ExecuteStatementInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ExecuteStatementOutput, error)
}

type DDBClient struct {
	DynamoDbClient *dynamodb.Client
}

func (r *DDBClient) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	return r.DynamoDbClient.GetItem(ctx, params, optFns...)
}

func (r *DDBClient) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	return r.DynamoDbClient.PutItem(ctx, params, optFns...)
}

func (r *DDBClient) Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	return r.DynamoDbClient.Scan(ctx, params, optFns...)
}
func (r *DDBClient) DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error) {
	return r.DynamoDbClient.DeleteItem(ctx, params, optFns...)
}

func (r *DDBClient) UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {
	return r.DynamoDbClient.UpdateItem(ctx, params, optFns...)
}
func (r *DDBClient) ExecuteStatement(ctx context.Context, params *dynamodb.ExecuteStatementInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ExecuteStatementOutput, error) {
	return r.DynamoDbClient.ExecuteStatement(ctx, params, optFns...)
}

type TableScanResult struct {
	TableName string
	Items     []map[string]types.AttributeValue
	Err       error
}

func GetHistory(ctx context.Context, ddbClient DDBClientiface, tableName, date string) (*dynamodb.ExecuteStatementOutput, error) {
	query := fmt.Sprintf("select * from \"%s\" where begins_with(venueDate, ?)", tableName)
	params := []types.AttributeValue{
		&types.AttributeValueMemberS{Value: date},
	}
	result, err := ddbClient.ExecuteStatement(ctx, &dynamodb.ExecuteStatementInput{
		Statement:  aws.String(query),
		Parameters: params,
	})

	if err != nil {
		return nil, err
	}
	return result, nil

}

func ScanTable(ctx context.Context, ddbClient DDBClientiface, tableName string) (*dynamodb.ScanOutput, error) {
	// Scan 실행
	resp, err := ddbClient.Scan(ctx, &dynamodb.ScanInput{
		TableName: &tableName,
	})

	if err != nil {
		return nil, err
	}

	return resp, err
}

func DeletePendingItem(ctx context.Context, ddbClient DDBClientiface, key map[string]types.AttributeValue) error {
	tableName := "pending_reservation"

	_, err := ddbClient.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &tableName,
		Key:       key,
	})

	if err != nil {
		log.Errorln("err", err)
	}

	return err
}
func DeleteReservationItem(ctx context.Context, ddbClient DDBClientiface, key map[string]types.AttributeValue) error {
	tableName := "current_reservation"

	_, err := ddbClient.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &tableName,
		Key:       key,
	})

	return err
}

func IsItemExist(ctx context.Context, ddbClient DDBClientiface, tableName string, key map[string]types.AttributeValue) (map[string]types.AttributeValue, error) {
	result, err := ddbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &tableName,
		Key:       key,
	})

	if err != nil {
		log.Errorln("err", err)
		return nil, err
	}

	return result.Item, nil
}
func GetPendingItem(ctx context.Context, ddbClient DDBClientiface, key map[string]types.AttributeValue) (map[string]types.AttributeValue, error) {
	tableName := "pending_reservation"

	result, err := ddbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &tableName,
		Key:       key,
	})

	if err != nil {
		log.Errorln("err", err)
	}

	return result.Item, err
}

func AcceptReservation(ctx context.Context, ddbClient DDBClientiface, reservationInfo map[string]types.AttributeValue) error {
	tableName := "current_reservation"

	_, err := ddbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"reservationId": reservationInfo["requestId"],
			"category":      reservationInfo["category"],
			"companion":     reservationInfo["companion"],
			"email":         reservationInfo["email"],
			"name":          reservationInfo["name"],
			"purpose":       reservationInfo["purpose"],
			"studentId":     reservationInfo["studentId"],
			"time":          reservationInfo["time"],
			"venueDate":     reservationInfo["venueDate"],
		},
	})
	if err != nil {
		log.Errorln("err", err)
	}
	return err
}

func ChangeReservationValues(ctx context.Context, ddbClient DDBClientiface, key map[string]types.AttributeValue, values map[string]types.AttributeValue) error {
	tableName := "current_reservation"

	date := values["Date"].(*types.AttributeValueMemberS).Value
	venue := values["Venue"].(*types.AttributeValueMemberS).Value
	time := values["ChangeTime"].(*types.AttributeValueMemberL)

	log.Info(time, venue)
	venueDate := fmt.Sprintf("%s#%s", date, venue)
	// 업데이트할 속성 정의
	update := expression.
		Set(expression.Name("time"), expression.Value(time)).
		Set(expression.Name("venueDate"), expression.Value(venueDate))

	// 업데이트 표현식 생성
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		log.Printf("Couldn't build expression for update. Error: %v\n", err)
		return err
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
		log.Printf("Couldn't update item. Error: %v\n", err)
		return err
	}

	return err
}
