package main

import (
	"log"
	"os"

	"hotelifoz/cmd/app"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	route := app.StartApp()

	log.Println("Starting the application...")
	route.Start(os.Getenv("APP_PORT"))
}
