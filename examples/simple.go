package main

import (
	"fmt"
	"github.com/lalabuy948/genvutils"
)

type config struct {
	ServerPort string `genv:"SERVER_PORT, 8080"` // takes 8080 if SERVER_PORT env value is null
}

func main()  {
	if genvutils.IsDevelopment() {
		fmt.Println("This will be executed if development")
	}

	var cfg config
	genvutils.Parse(&cfg)

	fmt.Println("server port is:", cfg.ServerPort) // -> server port is: 8080
}
