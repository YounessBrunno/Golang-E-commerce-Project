package env

import "os"

func GetEnv(key, fallbackValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallbackValue
}