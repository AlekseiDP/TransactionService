package main

import (
	"TransactionService/internal/composites"
	"TransactionService/internal/domain/account"
	"TransactionService/internal/domain/user"
	"log"
)

func main() {
	log.Print("Initializing postgres composite")
	postgresComposite, err := composites.NewPostgresComposite()
	if err != nil {
		log.Fatal("Error Initializing postgres composite")
	}

	postgresComposite.DB.AutoMigrate(&account.Account{})
	postgresComposite.DB.AutoMigrate(&user.User{})
}
