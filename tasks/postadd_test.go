package tasks

import (
	"backend-worktask/tests"
	"net/http"
	"strings"
	"testing"
)

func TestAddHandler_MissingB(t *testing.T) {
	// Create a request body with only 'a' provided
	requestBody := `{"a": 3}`

	// Create a request
	rr := tests.MakeTestRequest("POST", "/task/add", requestBody, PostAddHandler, t)

	// Check the status code
	tests.CheckStatusCode(rr.Code, http.StatusBadRequest, t)
}

func TestAddHandler_InvalidA(t *testing.T) {
	// Create a request body with invalid 'a' (string instead of integer)
	requestBody := `{"a": "invalid", "b": 5}`

	// Create a request
	rr := tests.MakeTestRequest("POST", "/task/add", requestBody, PostAddHandler, t)

	// Check the status code
	tests.CheckStatusCode(rr.Code, http.StatusBadRequest, t)
}

func TestAddHandler_ValidInput(t *testing.T) {
	// Create a request body with valid arguments
	requestBody := `{"a": 3, "b": 5}`

	// Create a request
	rr := tests.MakeTestRequest("POST", "/task/add", requestBody, PostAddHandler, t)

	// Check the status code
	tests.CheckStatusCode(rr.Code, http.StatusOK, t)

	// Check the response body
	expected := `{"result":8}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
