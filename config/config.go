package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	Debug bool
	//ServerPort int64
	ServerConfig 
}

type ServerConfig struct {
	ServerPort string
	ReadTimeout int64
	WriteTimeout int64
	MaxHeaderBytes int64
}

func GetEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
	return &Env{
		Debug: parseBool(os.Getenv("DEBUG")),
		ServerConfig: ServerConfig{
			ServerPort: os.Getenv("SERVER_PORT"),
			ReadTimeout: parseInt64(os.Getenv("READ_TIMEOUT"), 10),
			WriteTimeout: parseInt64(os.Getenv("WRITE_TIMEOUT"), 10),
			MaxHeaderBytes: parseInt64(os.Getenv("MAX_HEADER_BYTES"), 1 << 20),
		},
	}
}

func parseBool(key string) bool {
	boolValue, err := strconv.ParseBool(key)
	if err != nil {
		return true
	}
	return boolValue
} 

func parseInt64(key string, defaultValue int64) int64 {
	intValue, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return defaultValue
	}
	return intValue
}