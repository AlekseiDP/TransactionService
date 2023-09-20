package main

import (
	"TransactionService/internal/composites"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// entry point
	log.Print("Initializing .env file")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Initializing .env file")
	}

	log.Print("Initializing postgres composite")
	postgresComposite, err := composites.NewPostgresComposite()
	if err != nil {
		log.Fatal("Error Initializing postgres composite")
	}

	r := gin.Default()
	log.Print("Initializing account composite")
	accountComposite, err := composites.NewAccountComposite(postgresComposite)
	if err != nil {
		log.Fatal("Error Initializing account composite")
	}
	accountComposite.Handler.Register(r)

	if err := r.Run(); err != nil {
		log.Fatal("Error starting server")
	}
}
