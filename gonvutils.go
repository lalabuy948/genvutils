// Package gonvutils provides useful environment operations
package gonvutils

import "os"

//IsProduction checks if ENVIRONMENT value is equal to "PROD".
func IsProduction() bool {
	if os.Getenv("ENVIRONMENT") == "PROD" {
		return true
	}
	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		return true
	}
	return false
}

//IsDevelopment checks if ENVIRONMENT value is equal to "DEV".
func IsDevelopment() bool {
	if os.Getenv("ENVIRONMENT") == "DEV" {
		return true
	}
	if os.Getenv("ENVIRONMENT") == "DEVELOPMENT" {
		return true
	}
	return false
}

//IsTesting checks if ENVIRONMENT value is equal to "TEST".
func IsTesting() bool {
	if os.Getenv("ENVIRONMENT") == "TEST" {
		return true
	}
	if os.Getenv("ENVIRONMENT") == "TESTING" {
		return true
	}
	return false
}

//GetEnv extracts env value by provided key, otherwise falls back to second function argument.
func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
