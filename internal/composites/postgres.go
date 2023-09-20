package composites

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type PostgresComposite struct {
	DB *gorm.DB
}

func NewPostgresComposite() (*PostgresComposite, error) {
	dsn := os.Getenv("CONNECTION_STRING")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &PostgresComposite{
		DB: DB,
	}, nil
}
