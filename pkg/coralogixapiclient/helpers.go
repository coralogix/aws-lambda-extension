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
	var severity int

	message = strings.ToLower(message)

	switch {
	case strings.Contains(message, "debug"):
		severity = 1
	case strings.Contains(message, "verbose"), strings.Contains(message, "trace"):
		severity = 2
	case strings.Contains(message, "warning"), strings.Contains(message, "warn"):
		severity = 4
	case strings.Contains(message, "error"), strings.Contains(message, "exception"):
		severity = 5
	case strings.Contains(message, "fatal"), strings.Contains(message, "critical"):
		severity = 6
	default:
		severity = 3
	}

	return severity
}
