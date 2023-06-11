package config

import (
	"log"

	"github.com/caarlos0/env/v8"
)

type Config struct {
	MongoUri            string `env:"MONGO_URI,required"`
	MongoDB             string `env:"MONGO_DB,required"`
	MongoUserCollection string `env:"MONGO_USER_COLLECTION,required"`
}

func InitConfig() *Config {

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}
