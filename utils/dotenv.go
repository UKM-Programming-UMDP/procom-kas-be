package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	if godotenv.Load(".env") != nil { // for dev
		if godotenv.Load("../../.env") != nil { // for test
			panic("Error loading .env file")
		}
	}
	return os.Getenv(key)
}
