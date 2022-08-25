package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"green.env.com/auth/util"
)

var config *Config

type Config struct {
	Port        string `envconfig:"PORT"`
	Stage       string `envconfig:"STAGE"`
	ServiceHost string `envconfig:"SERVICE_HOST"`
	MySQL       struct {
		Host           string `envconfig:"DB_HOST"`
		Port           string `envconfig:"DB_PORT"`
		User           string `envconfig:"DB_USER"`
		Pass           string `envconfig:"DB_PASS"`
		DBName         string `envconfig:"DB_NAME"`
		DBMaxIdleConns int    `envconfig:"DB_MAX_IDLE_CONNS"`
		DBMaxOpenConns int    `envconfig:"DB_MAX_OPEN_CONNS"`
		CountRetryTx   int    `envconfig:"DB_TX_RETRY_COUNT"`
	}
}

func init() {
	config = &Config{}

	_ = godotenv.Load()

	err := envconfig.Process("", config)
	if err != nil {
		err = errors.Wrap(err, "Failed to decode config env")
		util.GetLogger().Fatal(err.Error())
	}
}

func GetConfig() *Config {
	return config
}
