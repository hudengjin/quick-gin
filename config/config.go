package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	Debug bool

	ServerPort int64
}

func GetEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
	return &Env{
		Debug: parseBool(os.Getenv("DEBUG")),
		ServerPort: parseInt(os.Getenv("SERVER_PORT")),
	}
}

func parseBool(key string) bool {
	boolValue, err := strconv.ParseBool(key)
	if err != nil {
		return true
	}
	return boolValue
} 

func parseInt(key string) int64 {
	intValue, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return 8080
	}
	return intValue
}