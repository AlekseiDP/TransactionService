package main

import (
	"TransactionService/cmd/initializers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	initializers.LoadEnv()
	initializers.ConnectToDb()

	r := gin.Default()
	initializers.RegisterHandlers(r)

	if err := r.Run(); err != nil {
		log.Fatal("Error starting server")
	}
}
