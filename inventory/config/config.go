package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func GetEnv() string {
	return getEnvironmentValue("ENV")
}

func GetDatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		getEnvironmentValue("DB_USERNAME"),
		getEnvironmentValue("DB_PASSWORD"),
		getEnvironmentValue("DB_HOST"),
		getEnvironmentValue("DB_PORT"),
		getEnvironmentValue("DB_DATABASE"),
		getEnvironmentValue("DB_SCHEMA"),
	)
}

func GetApplicationPort() int {
	portStr := getEnvironmentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}

	return port
}

func getEnvironmentValue(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}

	return os.Getenv(key)
}
