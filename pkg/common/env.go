package common

import (
	"os"
)

var (
	staticFilesPath      string
	authFilesPath        string
	uploadsDirectoryPath string
	logLevel             string
)

// GetEnv checks for existing environment variables
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
