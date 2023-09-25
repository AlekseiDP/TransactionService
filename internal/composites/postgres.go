package composites

import (
	"TransactionService/internal/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresComposite Структура для регистрации ORM
type PostgresComposite struct {
	DB *gorm.DB
}

func NewPostgresComposite() (*PostgresComposite, error) {
	databaseConfig := config.GetDatabaseConfig()
	connectionString := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", databaseConfig.Db.Host, databaseConfig.Db.User, databaseConfig.Db.Password, databaseConfig.Db.Database, databaseConfig.Db.Port)
	DB, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &PostgresComposite{
		DB: DB,
	}, nil
}
