package coralogixapiclient

import (
	"os"
	"strings"
)

// GetEnv extract environment variable or default value
func GetEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

// GetSeverityLevel extract serverity from log message
func GetSeverityLevel(message string) int {
	severity := 3
	if strings.Contains(message, "Warning") || strings.Contains(message, "warn") {
		severity = 4
	}
	if strings.Contains(message, "Error") || strings.Contains(message, "Exception") {
		severity = 5
	}
	return severity
}
