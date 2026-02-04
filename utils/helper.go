package utils

import (
	"os"
	"strconv"
)

func GetEnvAsInt(key string, fallback int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return fallback
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return fallback
	}

	return value
}