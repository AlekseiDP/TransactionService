package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type DatabaseConfig struct {
	Db struct {
		Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
		User     string `yaml:"user" env:"DB_USER" env-default:"pgsu"`
		Password string `yaml:"password" env:"DB_PASSWORD" env-default:"Qwerty123!"`
		Database string `yaml:"database" env:"DB_DATABASE" env-default:"TransactionService"`
		Port     string `yaml:"port" env:"DB_PORT" env-default:"5433"`
	} `yaml:"db"`
}

var databaseConfig *DatabaseConfig
var once sync.Once

func GetDatabaseConfig() *DatabaseConfig {
	once.Do(func() {
		databaseConfig = &DatabaseConfig{}
		if err := cleanenv.ReadConfig("config.yml", databaseConfig); err != nil {
			help, _ := cleanenv.GetDescription(databaseConfig, nil)
			log.Print(help)
			log.Fatal(err)
		}
	})

	return databaseConfig
}
