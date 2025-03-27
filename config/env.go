package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBHost                 string
	DBPort                 string
	DBAddress              string
	DBName                 string
	DBUser                 string
	DBPassword             string
	APP_PORT               string
	JWTSecret              string
	JWTExpirationInSeconds string
}

var Envs = initConfig()

func initConfig() Config {
	_ = godotenv.Load()
	return Config{
		DBUser:                 os.Getenv("DB_USER"),
		DBPassword:             os.Getenv("DB_PASSWORD"),
		DBAddress:              fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName:                 os.Getenv("DB_NAME"),
		APP_PORT:               os.Getenv("APP_PORT"),
		JWTSecret:              os.Getenv("JWT_SECRET"),
		JWTExpirationInSeconds: os.Getenv("JWT_EXPIRATION_IN_SECONDS"),
	}
}
