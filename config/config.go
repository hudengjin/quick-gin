package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Env 环境配置参数结构体
type Env struct {
	AppName string
	Debug bool
	LogPath string
	LogMaxSize int
	LogMaxAge int
	LogMaxBackups int
	LogIsCompress bool
	//ServerPort int64

	ServerConfig 
}

// ServerConfig 服务器配置 
type ServerConfig struct {
	ServerPort string
	ReadTimeout int64
	WriteTimeout int64
	MaxHeaderBytes int
}

// GetEnv 获取环境配置
func GetEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
	return &Env{
		AppName: os.Getenv("APP_NAME"),
		Debug: parseBool(os.Getenv("DEBUG"), false),
		LogPath: os.Getenv("LOG_PATH"),
		LogMaxSize: parseInt(os.Getenv("LOG_MAX_SIZE"), 128),
		LogMaxAge: parseInt(os.Getenv("LOG_MAX_AGE"), 7),
		LogMaxBackups: parseInt(os.Getenv("LOG_MAX_BACKUPS"), 2),
		LogIsCompress: parseBool(os.Getenv("LOG_IS_COMPRESS"), true),
		ServerConfig: ServerConfig{
			ServerPort: os.Getenv("SERVER_PORT"),
			ReadTimeout: parseInt64(os.Getenv("READ_TIMEOUT"), 10),
			WriteTimeout: parseInt64(os.Getenv("WRITE_TIMEOUT"), 10),
			MaxHeaderBytes: parseInt(os.Getenv("MAX_HEADER_BYTES"), 1 << 20),
		},
	}
}

func parseBool(key string, defaultValue bool) bool {
	boolValue, err := strconv.ParseBool(key)
	if err != nil {
		return defaultValue
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

func parseInt(key string, defaultValue int) int {
	intValue, err := strconv.Atoi(key)
	if err != nil {
		return defaultValue
	}
	return intValue
}