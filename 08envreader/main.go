package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./../.env")

	if err != nil {

		log.Fatal("Error loading .env file")
	}

	appPort := os.Getenv("APP_PORT")

	fmt.Println("App Port:", appPort)
}
