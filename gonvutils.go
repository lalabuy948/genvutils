// Package genvutils provides useful environment operations
package genvutils

import (
	"os"
	"reflect"
	"strconv"
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

//Parse function will parse given pointer to struct and fill it with env values.
//
//  type serverConfig struct {
//      ServerPort string `genv:"SERVER_PORT,8080"`
//      MongoUrl   string `genv:"MONGO_URL,mongodb://localhost:27017"`
//  }
//
// Here is an example of struct. Good reading about reflect https://github.com/a8m/reflect-examples
func Parse(income interface{}) error {
	t := reflect.TypeOf(income).Elem()
	v := reflect.ValueOf(income).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.FieldByName(field.Name)

		tag := field.Tag.Get("genv")
		tagS := strings.Split(tag, ",")

		if value.CanSet() && value.IsValid() {
			switch len(tagS) {
			case 0:
				break
			default:
				envVarValue := strings.TrimSpace(GetEnv(tagS[0], strings.Join(tagS[1:], ",")))
				switch value.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					envVarValueF, _ := strconv.ParseInt(envVarValue, 10, 64)
					value.SetInt(envVarValueF)
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					envVarValueF, _ := strconv.ParseUint(envVarValue, 10, 64)
					value.SetUint(envVarValueF)
				case reflect.Float32, reflect.Float64:
					envVarValueF, _ := strconv.ParseFloat(envVarValue, 64)
					value.SetFloat(envVarValueF)
				case reflect.String:
					value.SetString(envVarValue)
				case reflect.Bool:
					envVarValueF, _ := strconv.ParseBool(envVarValue)
					value.SetBool(envVarValueF)
				}
			}
		}
	}

	return nil
}
