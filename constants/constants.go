package constants

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}
