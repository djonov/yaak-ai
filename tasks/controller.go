package tasks

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type PostAddRequestBody struct {
	A int `json:"a" validate:"required,number"`
	B int `json:"b" validate:"required,number"`
}

type PostAddResponseBody struct {
	Result int `json:"result"`
}

type GetDriveResponseBody struct {
	Id    int         `json:"id"`
	Route interface{} `json:"route"`
}

func PostAddHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[Controller.PostAddHandler] Status: started")
	var requestBody PostAddRequestBody

	// Decode the JSON request body
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Printf("[Controller.PostAddHandler] Status: failed, Error: Malformed JSON")
		http.Error(w, "Malformed JSON", http.StatusBadRequest)
		return
	}

	// Validate
	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		http.Error(w, validationErrors.Error(), http.StatusBadRequest)
		log.Printf("[Controller.PostAddHandler] Status: failed, Error: %s", validationErrors.Error())
		return
	}

	// Calculate the sum
	sum := Add(requestBody.A, requestBody.B)

	// Create the response body
	responseBody := PostAddResponseBody{Result: sum}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseBody)
	log.Printf("[Controller.PostAddHandler] Status: success")
}

func GetDriveHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[Controller.GetDriveHandler] Status: started")

	// Extract id path param from URL
	idStr := strings.TrimPrefix(r.URL.Path, "/task/drive/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[Controller.GetDriveHandler] Status: failed, Error: Invalid ID")
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Get external service response
	routeData := GetDriveById()

	// Create the response body
	responseBody := GetDriveResponseBody{Id: id, Route: routeData}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseBody)
	log.Printf("[Controller.GetDriveHandler] Status: success")
}
