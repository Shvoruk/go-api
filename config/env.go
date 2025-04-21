package config

import (
	"github.com/joho/godotenv"
	"os"
)

// Config holds application settings
type Config struct {
	DBAddress  string
	DBName     string
	DBUser     string
	DBPassword string
	JWTSecret  string
}

var Envs = initConfig()

func initConfig() Config {
	_ = godotenv.Load()
	return Config{
		DBUser:     os.Getenv("MYSQL_USER"),
		DBPassword: os.Getenv("MYSQL_PASSWORD"),
		DBAddress:  os.Getenv("MYSQL_HOST"),
		DBName:     os.Getenv("MYSQL_DATABASE"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}
}
