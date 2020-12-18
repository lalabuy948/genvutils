package main

import (
	"fmt"
	"log"

	"github.com/lalabuy948/genvutils"
)

type config struct {
	ServerPort string `genv:"SERVER_PORT, 8080"` // takes 8080 if SERVER_PORT env value is null
}

func main() {
	// load dotenv file
	err := genvutils.Load()
	if err != genvutils.ErrDotenvNotFound {
		log.Fatal(err)
	}

	// check if development for some reason ¯\_(ツ)_/¯
	if genvutils.IsDevelopment() {
		fmt.Println("This will be executed if development")
	}

	// dump env variables to your struct
	var cfg config
	genvutils.Parse(&cfg)

	fmt.Println("server port is:", cfg.ServerPort) // -> server port is: 8080
}
