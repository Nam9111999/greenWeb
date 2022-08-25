package auth

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"green.env.com/auth/util"
)

var config *Config

type Config struct {
	Port        string `envconfig:"PORT"`
	ServiceHost string `envconfig:"SERVICE_HOST"`
	Stage       string `envconfig:"STAGE"`
}

func init() {
	config = &Config{}

	_ = godotenv.Load()

	err := envconfig.Process("", config)
	if err != nil {
		util.GetLogger().Fatal(err.Error())
	}
}
