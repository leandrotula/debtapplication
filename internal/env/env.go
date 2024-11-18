package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func GetString(key, defaultValue string) string {
	loadEnv()

	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}
	return result
}

func GetInt(key string, defaultValue int) int {
	loadEnv()
	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}

	convertedValue, err := strconv.Atoi(result)
	if err != nil {
		panic(err)
	}
	return convertedValue
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
