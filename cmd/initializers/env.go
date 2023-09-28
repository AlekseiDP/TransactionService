package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	log.Print("Initializing .env file")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Initializing .env file")
	}
}
