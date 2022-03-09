package main

import (
	"github.com/joho/godotenv"
	"github/atakanteko/bookstore_users-api/app"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app.StartApplication()
}
