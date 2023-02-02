package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Env            string
	GIN_MODE       string `envconfig:"GIN_MODE" default:"release"`
	LISTENING_PORT string `envconfig:"LISTENING_PORT" default:"3000"`
	PROXY_URL      string `envconfig:"PROXY_URL"`
}

func Load() Config {
	var config Config
	ENV, ok := os.LookupEnv("ENV")

	if !ok {
		// Default value for ENV.
		ENV = "dev"
	}

	err := godotenv.Load("./.env")
	if err != nil {
		logrus.Warn("Can't load env file")
	}

	envconfig.MustProcess("", &config)
	config.Env = ENV
	return config
}