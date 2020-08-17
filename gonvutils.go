// Package genvutils provides useful environment operations
package genvutils

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

//IsProduction checks if ENVIRONMENT value is equal to "PROD".
func IsProduction() bool {
	if os.Getenv("ENVIRONMENT") == "PROD" {
		return true
	}
	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		return true
	}
	if os.Getenv("APP_ENV") == "PROD" {
		return true
	}
	if os.Getenv("APP_ENV") == "PRODUCTION" {
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
	if os.Getenv("APP_ENV") == "DEV" {
		return true
	}
	if os.Getenv("APP_ENV") == "DEVELOPMENT" {
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
	if os.Getenv("APP_ENV") == "TEST" {
		return true
	}
	if os.Getenv("APP_ENV") == "TESTING" {
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

func Parse(i interface{}) interface{} {
	t := reflect.TypeOf(i)

	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get("genv")
		tagS := strings.Split(tag, ",")
		tagName := tagS[0]
		defValue := tagS[1]

		fmt.Printf("envkey: %v defvalue: %v \n", tagName, defValue)
		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}

	return t
}
