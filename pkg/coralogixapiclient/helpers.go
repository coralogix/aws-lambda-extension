package coralogixapiclient

import (
	"os"
)

// GetEnv extract environment variable or default value
func GetEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
