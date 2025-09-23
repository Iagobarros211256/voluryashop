package configs

import (
	"log"
	"os"
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
