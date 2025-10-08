package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseConfig *DatabaseConfig
	LogConfig      *LogConfig
	AppConfig      *AppConfig
}

type AppConfig struct {
	Port string
}

type DatabaseConfig struct {
	Url string
}

type LogConfig struct {
	Level  int
	Format string
}

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Print("error loading .env file")
		return
	}
	log.Println("config file loaded")
}

func getString(key, defaultString string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultString
	}
	return value
}
func getBool(key string, defaultBool bool) bool {
	valueStr := os.Getenv(key)
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return defaultBool
	}
	return value
}

func getInt(key string, defaultInt int) int {
	valueStr := os.Getenv(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultInt
	}
	return value
}

func NewConfig() *Config {
	return &Config{
		DatabaseConfig: &DatabaseConfig{
			Url: getString("DATABASE_URL", ""),
		},
		LogConfig: &LogConfig{
			Level:  getInt("LOG_LEVEL", 0),
			Format: getString("LOG_FORMAT", "json"),
		},
		AppConfig: &AppConfig{
			Port: getString("APP_PORT", ":3000"),
		},
	}
}
