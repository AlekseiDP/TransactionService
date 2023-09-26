package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type JwtConfig struct {
	Jwt struct {
		Issuer               string `yaml:"issuer" env-default:"http://localhost:3000"`
		Audience             string `yaml:"audience" env-default:"TransactionService"`
		AccessTokenLifetime  int    `yaml:"accessTokenLifetime" env-default:"900"`
		RefreshTokenLifetime int    `yaml:"refreshTokenLifetime" env-default:"259200"`
		Secret               string `yaml:"secret" env:"JWT_SECRET" env-default:"MySuperSecretKey"`
	} `yaml:"jwt"`
}

var jwtConfig *JwtConfig

func GetJwtConfig() *JwtConfig {
	once.Do(func() {
		jwtConfig = &JwtConfig{}
		if err := cleanenv.ReadConfig("config.yml", databaseConfig); err != nil {
			help, _ := cleanenv.GetDescription(databaseConfig, nil)
			log.Print(help)
			log.Fatal(err)
		}
	})

	return jwtConfig
}
