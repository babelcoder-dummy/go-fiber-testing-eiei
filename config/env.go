package config

import (
	"os"

	"github.com/joho/godotenv"
)

type secret struct {
	AccessToken  string
	RefreshToken string
}

type expiration struct {
	AccessToken  string
	RefreshToken string
}

type env struct {
	Host        string
	Port        string
	DatabaseURL string
	Secret      secret
	Expiration  expiration
}

var Env env

func InitEnv() {
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "" {
		appEnv = "development"
	}

	godotenv.Load(".env." + appEnv)
	godotenv.Load()

	Env = env{
		Host:        os.Getenv("HOST"),
		Port:        os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Secret: secret{
			AccessToken:  os.Getenv("ACCESS_TOKEN_SECRET_KEY"),
			RefreshToken: os.Getenv("REFRESH_TOKEN_SECRET_KEY"),
		},
		Expiration: expiration{
			AccessToken:  os.Getenv("ACCESS_TOKEN_EXPIRES_IN"),
			RefreshToken: os.Getenv("REFRESH_TOKEN_EXPIRES_IN"),
		},
	}
}
