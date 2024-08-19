package util

import "os"

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
