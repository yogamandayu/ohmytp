package util

import (
	"os"
	"strconv"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func GetEnvAsBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value != "" {
		return value == "true"
	}
	return defaultValue
}

func GetEnvAsInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value != "" {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return intValue
	}
	return defaultValue
}
