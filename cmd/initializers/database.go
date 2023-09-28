package initializers

import (
	"TransactionService/internal/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDb() {
	log.Print("Initializing db connection")
	var err error
	databaseConfig := config.GetDatabaseConfig()
	connectionString := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", databaseConfig.Db.Host, databaseConfig.Db.User, databaseConfig.Db.Password, databaseConfig.Db.Database, databaseConfig.Db.Port)
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal(fmt.Sprintf("Error initializing db connection: %v", err.Error()))
	}
}
