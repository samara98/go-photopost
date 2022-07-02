package main

import (
	"log"

	"github.com/joho/godotenv"
)

//go:generate go run github.com/google/wire/cmd/wire

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Default().Println(err.Error())
	}

	app := InitApp()

	app.Start()
}
