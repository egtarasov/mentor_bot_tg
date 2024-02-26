package main

import (
	"github.com/joho/godotenv"
	"log"
	"telegrambot_new_emploee/internal/api"
)

func main() {
	// TODO better configuration
	err := godotenv.Load("/Users/egtarasov/University/Projects/telegrambot_ne_employe/deploy/.env")
	if err != nil {
		log.Fatal(err)
	}

	//Play()
	api.Run()
}
