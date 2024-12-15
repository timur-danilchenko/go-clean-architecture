package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/lpernett/godotenv"
)

const envPath = ".env"

type Config struct {
	PORT    string `default:8080`
	DB_USER string
	DB_PASS string
	DB_HOST string `default:localhost`
	DB_PORT string `default:5432`
	DB_NAME string `default:postgres`
}

func Setup() (conf Config, err error) {
	_ = godotenv.Load(envPath)

	if err = envconfig.Process("", &conf); err != nil {
		return
	}

	return
}
