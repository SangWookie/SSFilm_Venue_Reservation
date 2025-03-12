package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"request_manager/actions"
	"request_manager/response"
	"sync"

	"github.com/aws/aws-lambda-go/events"
)

func GetReservations(ctx context.Context, request events.APIGatewayV2HTTPRequest, ddbClient actions.DDBClientiface) (events.APIGatewayV2HTTPResponse, error) {
	tables := []string{"current_reservation", "pending_reservation"}

	// go routine 사용해서 병렬 처리1
	var wg sync.WaitGroup
	results := make(chan actions.TableScanResult, len(tables))

	for _, tableName := range tables {
		wg.Add(1)
		go actions.ScanTable(ctx, ddbClient, tableName, &wg, results)
	}

	// Goroutine 완료 대기
	wg.Wait()
	close(results)

	var scanResults map[string]actions.TableScanResult
	scanResults = make(map[string]actions.TableScanResult)

	for result := range results {
		if result.Err != nil {
			log.Printf("Error scanning table %s: %v", result.TableName, result.Err)
		} else {
			fmt.Printf("Scanned %d items from %s\n", len(result.Items), result.TableName)
		}

		switch result.TableName {
		case "current_reservation":
			scanResults["reservation history"] = result
		case "pending_reservation":
			scanResults["pending list"] = result
		}
	}
	return response.APIGatewayResponseOK(scanResults, http.StatusOK), nil
}
