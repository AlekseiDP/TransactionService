package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
	"time"
)

type JwtConfig struct {
	Jwt struct {
		Issuer               string        `yaml:"issuer" env-default:"http://localhost:3000"`
		Audience             string        `yaml:"audience" env-default:"TransactionService"`
		AccessTokenLifetime  time.Duration `yaml:"accessTokenLifetime" env-default:"900"`
		RefreshTokenLifetime time.Duration `yaml:"refreshTokenLifetime" env-default:"259200"`
		Secret               string        `yaml:"secret" env:"JWT_SECRET" env-default:"MySuperSecretKey"`
	} `yaml:"jwt"`
}

var jwtConfig *JwtConfig
var once2 sync.Once

func GetJwtConfig() *JwtConfig {
	once2.Do(func() {
		jwtConfig = &JwtConfig{}
		if err := cleanenv.ReadConfig("config.yml", jwtConfig); err != nil {
			help, _ := cleanenv.GetDescription(jwtConfig, nil)
			log.Print(help)
			log.Fatal(err)
		}
	})

	return jwtConfig
}
