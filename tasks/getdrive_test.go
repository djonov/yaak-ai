package tasks

import (
	"backend-worktask/config"
	"backend-worktask/tests"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"

	"github.com/h2non/gock"
)

func TestGetDriveHandler_Success(t *testing.T) {
	// Set environment variables
	tests.SetEnvVariables()

	// Get config to get env variables
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("[Tests.TestGetDriveHandler_Success] Info: error loading configuration: %v", err)
	}

	// Disable default HTTP interceptors
	defer gock.Off()

	// Define mock external service response body
	externalServiceJSON := map[string]interface{}{
		"value": "fixed",
	}
	id := 1

	// Mock the external HTTP endpoint twice.
	// First it returns 429 and then it returns 200
	gock.New(cfg.GET_DRIVE_BY_ID_URL).
		Get("/").
		Reply(429)

	gock.New(cfg.GET_DRIVE_BY_ID_URL).
		Get("/").
		Reply(200).
		JSON(externalServiceJSON)

	// Consume the endpoint for testing
	rr := tests.MakeTestRequest("GET", fmt.Sprintf("/task/drive/%d", id), "", GetDriveHandler, t)

	// Check the status code
	tests.CheckStatusCode(rr.Code, http.StatusOK, t)

	// Prepare expected object
	expected := GetDriveResponseBody{
		Id:    id,
		Route: externalServiceJSON,
	}

	// Decode response body
	var actual GetDriveResponseBody
	if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	// Check if the response body is the same
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("handler returned unexpected body: got %+v want %+v",
			actual, expected)
	}
}
