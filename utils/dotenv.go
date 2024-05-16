package utils

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	_, b, _, _ := runtime.Caller(0)
	rootPath := filepath.Join(filepath.Dir(b), "../")
	err := godotenv.Load(rootPath + "/.env")
	if err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(key)
}
