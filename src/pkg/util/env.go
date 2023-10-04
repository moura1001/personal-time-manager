package util

import "os"

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if key != "" {
		return value
	}
	return defaultValue
}
