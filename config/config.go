package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	QiscusAuthURL	string
	QiscusEmail		string
	QiscusPassword	string

	RedisHost		string
	RedisUser		string
	RedisPassword	string
	RedisPort		string
}

var AppConfig Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading from environment variables")
	}

	AppConfig = Config{
		QiscusAuthURL:		getEnv("QISCUS_AUTH_URL", ""),
		QiscusEmail:		getEnv("QISCUS_EMAIL", ""),
		QiscusPassword:		getEnv("QISCUS_PASSWORD", ""),

		RedisHost:			getEnv("REDIS_HOST", ""),
		RedisUser:			getEnv("REDIS_USER", ""),
		RedisPassword:		getEnv("REDIS_PASSWORD", ""),
		RedisPort:			getEnv("REDIS_PASSWORD", ""),
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
