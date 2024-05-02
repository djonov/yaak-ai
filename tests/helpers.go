package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MakeTestRequest(method string, path string, requestBody string, f http.HandlerFunc, t *testing.T) *httptest.ResponseRecorder {
	// Create a request
	var req *http.Request
	var err error

	if method == "GET" {
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
