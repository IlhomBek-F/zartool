package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	host                  string
	username              string
	password              string
	database              string
	DbPort                string
	Port                  string
	schema                string
	AccessTokenSecret     string
	AccessTokenExpiryHour string
}

func NewConfig() Config {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading env file: %v", err)
	}

	return Config{
		host:                  os.Getenv("DB_HOST"),
		DbPort:                os.Getenv("DB_PORT"),
		database:              os.Getenv("DB_DATABASE"),
		password:              os.Getenv("DB_PASSWORD"),
		username:              os.Getenv("DB_USERNAME"),
		schema:                os.Getenv("DB_SCHEMA"),
		Port:                  os.Getenv("Port"),
		AccessTokenSecret:     os.Getenv("ACCESS_TOKEN_SECRET"),
		AccessTokenExpiryHour: os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"),
	}
}
