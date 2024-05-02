package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func MakeTestRequest(method string, path string, requestBody string, f http.HandlerFunc, t *testing.T) *httptest.ResponseRecorder {
	// Create a request
	var req *http.Request
	var err error

	if requestBody == "" {
		// Create a request without a request body
		req, err = http.NewRequest(method, path, nil)
	} else {
		// Create a request with a request body
		req, err = http.NewRequest(method, path, bytes.NewBuffer([]byte(requestBody)))
	}

	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(f)

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)
	return rr
}

func CheckStatusCode(actual int, expected int, t *testing.T) {
	if actual != expected {
		t.Errorf("handler returned wrong status code: got %v want %v",
			actual, expected)
	}
}

func SetEnvVariables() {
	os.Setenv("GET_DRIVE_BY_ID_URL", "http://localhost")
	os.Setenv("EXTERNAL_SERVICE_API_KEY", "123abc")
	os.Setenv("EXTERNAL_SERVICE_WORK_TASK_AUTH", "123abc")
	os.Setenv("MAX_AMOUNT_OF_RETRIES", "10")
	os.Setenv("REPEAT_REQUEST_SECONDS_DELAY", "3")
}
