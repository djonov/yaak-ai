package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type RequestParams struct {
	Url     string
	Method  string
	Headers http.Header
}

func RequestRepeater(params RequestParams, retries int) interface{} {
	var responseBody interface{}

	for i := 0; i < retries; i++ {
		response := MakeRequest(params.Url, params.Method, params.Headers)
		if response.StatusCode == http.StatusOK {
			defer response.Body.Close()
			err := json.NewDecoder(response.Body).Decode(&responseBody)
			if err != nil {
				log.Fatal(err)
			}
			return responseBody
		} else if response.StatusCode != http.StatusTooManyRequests {
			return nil
		}
	}

	return nil
}

func MakeRequest(url string, method string, headers http.Header) *http.Response {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// add a custom header to the request
	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}
