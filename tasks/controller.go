package tasks

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type PostAddRequestBody struct {
	A int `json:"a" validate:"required,number"`
	B int `json:"b" validate:"required,number"`
}

type PostAddResponseBody struct {
	Result int `json:"result"`
}

type ParametersErrorResponse struct {
	Errors []string `json:"errors"`
}

func PostAddHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody PostAddRequestBody

	// Decode the JSON request body
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Malformed JSON", http.StatusBadRequest)
		return
	}

	// Validate
	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		http.Error(w, validationErrors.Error(), http.StatusBadRequest)
		return
	}

	// Calculate the sum
	sum := Add(requestBody.A, requestBody.B)

	// Create the response body
	responseBody := PostAddResponseBody{Result: sum}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseBody)
}
