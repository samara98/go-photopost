package main

import (
	"log"

	"github.com/joho/godotenv"
)

//go:generate go run github.com/google/wire/cmd/wire

func main() {
	logger := log.Default()

	err := godotenv.Load()
	if err != nil {
		logger.Println(err.Error())
	}

	server := InitServer()

	server.Start()
}
