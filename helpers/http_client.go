package helpers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type RequestParams struct {
	Url     string
	Method  string
	Headers http.Header
}

func RequestRepeater(params RequestParams, maxRetries int, delayInSeconds int32) interface{} {
	var responseBody interface{}

	for i := 0; i < maxRetries; i++ {
		response := MakeRequest(params.Url, params.Method, params.Headers)
		if response.StatusCode == http.StatusOK {
			defer response.Body.Close()
			err := json.NewDecoder(response.Body).Decode(&responseBody)
			if err != nil {
				log.Fatal(err)
			}
			return responseBody
		} else if response.StatusCode == http.StatusTooManyRequests {
			log.Printf("[HttpClient.RequesterRepeater] URL: %s, Method: %s, Info: sleeping for %d seconds because of a 429 status code", params.Url, params.Method, delayInSeconds)
			time.Sleep(time.Duration(rand.Int31n(delayInSeconds)) * time.Second)
		} else {
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
