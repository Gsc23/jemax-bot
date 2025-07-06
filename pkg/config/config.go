package config

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-simpler.org/env"
)

type Config struct {
	Whatsapp struct {
		VerifyToken string `env:"VERIFY_TOKEN" default:"jemax123"`
		AccessToken string `env:"ACCESS_TOKEN" default:"jemax123"`
		PhoneID string `env:"PHONE_ID" default:"jemax123"`
	} `env:"WHATSAPP"`
	Environment string `env:"ENVIRONMENT" default:"development"`
	ServerPort  int    `env:"SERVER_PORT" default:"8080"`
	Logger      struct {
		Level string `env:"LEVEL" default:"0"`
		Color bool   `env:"COLOR" default:"true"`
		Trace bool   `env:"TRACE" default:"true"`
	} `env:"LOGGER"`
	Database struct {
		Host     string `env:"HOST"`
		Port     int    `env:"PORT"`
		User     string `env:"USER"`
		Pass     string `env:"PASSWORD"`
		Database string `env:"DATABASE"`
	} `env:"DATABASE"`
}

func loadAsProd() {
	gin.SetMode("release")
}

func loadAsDev() {
	//
}

func Load(filenames ...string) *Config {
	godotenv.Load(filenames...)

	opts := &env.Options{NameSep: "_"}

	var config Config
	if err := env.Load(&config, opts); err != nil {
		panic(fmt.Sprintf("could not load from config: %v", err))
	}

	env, err := EnvironmentFromStr(config.Environment)
	if err != nil {
		panic(err)
	}

	switch env {
	case EnvProduction:
		loadAsProd()
	case EnvDevelop:
		loadAsDev()
	default:
		panic(fmt.Sprintf("invalid env: %s", config.Environment))
	}

	return &config
}
