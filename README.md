# genvutils  [![Go Report Card](https://goreportcard.com/badge/github.com/lalabuy948/genvutils)](https://goreportcard.com/report/github.com/lalabuy948/genvutils)  [![Build Status](https://github.com/lalabuy948/genvutils/workflows/build/badge.svg)](https://github.com/lalabuy948/genvutils/actions)

> Package genvutils provides useful environment operations

## Funcs

- `IsProduction`, `IsDevelopment` and `IsTesting` checks for `ENVIRONMENT` dot env value.
- `GetEnv` gets env value or fallback which goes as second function argument.
- `Parse` will fill given struct with env values or with fallbacks. (see examples folder)
```go
	type serverConfig struct {
		ServerPort string `genv:"SERVER_PORT,8080"`
		MongoUrl   string `genv:"MONGO_URL,mongodb://localhost:27017"`
	}
```

## Install

`go get -u github.com/lalabuy948/genvutils`
