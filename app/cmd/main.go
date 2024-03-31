package main

import (
	"log"
	"telegrambot_new_emploee/internal/api"
)

func main() {
	//err := godotenv.Load("/Users/egtarasov/University/Projects/telegrambot_ne_employe/deploy/.env")
	//if err != nil {
	//	log.Fatal(err)
	//}
	log.Printf("Now:")
	api.Run()
}
