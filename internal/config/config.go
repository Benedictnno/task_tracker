package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetDBURL(key string) string {
	return os.Getenv("DATABASE_URL")
}

func GetPort(key string) string {
	return os.Getenv("PORT")
}

func GetJWTSecret(key string) string {
	return os.Getenv("JWT_SECRET")
}
