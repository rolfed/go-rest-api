package env

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

func Load(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
