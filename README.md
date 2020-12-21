# genvutils  [![Go Report Card](https://goreportcard.com/badge/github.com/lalabuy948/genvutils)](https://goreportcard.com/report/github.com/lalabuy948/genvutils)  [![Build Status](https://github.com/lalabuy948/genvutils/workflows/build/badge.svg)](https://github.com/lalabuy948/genvutils/actions) [![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

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
- `Load` will load dotenv file. You can provide file name via argument otherwise it will load dot environment file by priority list.
First found - first load. Priority list:
```sh
.env.production.local
.env.test.local
.env.development.local
.env.production
.env.test
.env.development
.env.local !!! will override existing values.
.env
```

[example](examples/simple.go)

## Install

`go get -u github.com/lalabuy948/genvutils`

Donate

<a href="https://www.buymeacoffee.com/lalabuy" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-blue.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a>
