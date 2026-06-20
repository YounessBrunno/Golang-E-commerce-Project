package env

import "os"

func GetEnvKey(key, fallbackValue string) string {

	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallbackValue
}