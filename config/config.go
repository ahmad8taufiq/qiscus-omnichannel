package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	QiscusBaseURL	string
	QiscusApiURL	string

	QiscusEmail		string
	QiscusPassword	string

	QiscusAppID		string
	QiscusSecretKey	string

	RedisHost		string
	RedisUser		string
	RedisPassword	string
	RedisPort		string

	AdminToken		string
	SdkEmail		string
	SdkToken		string
}

var AppConfig Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading from environment variables")
	}

	AppConfig = Config{
		QiscusBaseURL:		getEnv("QISCUS_BASE_URL", ""),
		QiscusApiURL:		getEnv("QISCUS_API_URL", ""),

		QiscusEmail:		getEnv("QISCUS_EMAIL", ""),
		QiscusPassword:		getEnv("QISCUS_PASSWORD", ""),

		QiscusAppID:		getEnv("QISCUS_APP_ID", ""),
		QiscusSecretKey:	getEnv("QISCUS_SECRET_KEY", ""),

		RedisHost:			getEnv("REDIS_HOST", ""),
		RedisUser:			getEnv("REDIS_USER", ""),
		RedisPassword:		getEnv("REDIS_PASSWORD", ""),
		RedisPort:			getEnv("REDIS_PORT", ""),

		AdminToken:			"adminToken",
		SdkEmail:			"sdkEmail",
		SdkToken:			"sdkToken",
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
