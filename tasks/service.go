package tasks

import (
	"backend-worktask/config"
	"backend-worktask/helpers"
	"log"
	"net/http"
)

func Add(a int, b int) int {
	return a + b
}

func GetDriveById() interface{} {
	// Get config environment variables
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("[Service.GetDriveById] Info: error loading configuration: %v", err)
	}

	// Preparing headers
	headers := prepareGeometryServiceHeaders(cfg)

	// Prepare request params
	requestParams := helpers.RequestParams{Url: cfg.GET_DRIVE_BY_ID_URL, Method: "GET", Headers: headers}

	// Run HTTP request in repeater wrapper to avoid flaky service
	response := helpers.RequestRepeater(requestParams, cfg.MAX_AMOUNT_OF_RETRIES, int32(cfg.REPEAT_REQUEST_SECONDS_DELAY))

	return response
}

func prepareGeometryServiceHeaders(cfg *config.Config) http.Header {
	headers := make(http.Header)
	headers.Set("x-api-key", cfg.API_KEY)
	headers.Set("x-worktask-auth", cfg.WORK_TASK_AUTH)
	return headers
}
