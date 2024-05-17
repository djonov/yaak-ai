package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds the configuration values
type Config struct {
	PORT                         int
	GET_DRIVE_BY_ID_URL          string
	API_KEY                      string
	WORK_TASK_AUTH               string
	MAX_AMOUNT_OF_RETRIES        int
	REPEAT_REQUEST_SECONDS_DELAY int
}

// NewConfig creates a new Config instance with values read from environment variables
func GetConfig() (*Config, error) {
	maxAmountOfRetriesStr := os.Getenv("MAX_AMOUNT_OF_RETRIES")
	maxAmountOfRetries, err := strconv.Atoi(maxAmountOfRetriesStr)
	if err != nil {
		return nil, fmt.Errorf("invalid value for MAX_AMOUNT_OF_RETRIES: %v", err)
	}

	repeatRequestSecondsDelayStr := os.Getenv("REPEAT_REQUEST_SECONDS_DELAY")
	repeatRequestSecondsDelay, err := strconv.Atoi(repeatRequestSecondsDelayStr)
	if err != nil {
		return nil, fmt.Errorf("invalid value for REPEAT_REQUEST_SECONDS_DELAY: %v", err)
	}

	requiredEnvVars := []string{"GET_DRIVE_BY_ID_URL", "EXTERNAL_SERVICE_API_KEY", "EXTERNAL_SERVICE_WORK_TASK_AUTH"}
	envValues := make(map[string]string)
	for _, key := range requiredEnvVars {
		value := os.Getenv(key)
		if value == "" {
			return nil, fmt.Errorf("%s environment variable is not set", key)
		}
		envValues[key] = value
	}

	return &Config{
		GET_DRIVE_BY_ID_URL:          envValues["GET_DRIVE_BY_ID_URL"],
		API_KEY:                      envValues["EXTERNAL_SERVICE_API_KEY"],
		WORK_TASK_AUTH:               envValues["EXTERNAL_SERVICE_WORK_TASK_AUTH"],
		MAX_AMOUNT_OF_RETRIES:        maxAmountOfRetries,
		REPEAT_REQUEST_SECONDS_DELAY: repeatRequestSecondsDelay,
	}, nil
}
