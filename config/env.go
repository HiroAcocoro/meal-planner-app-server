package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/HiroAcocoro/meal-planner-app-server/internal/common/errors"
)

type Config struct {
	APIPort              string
	DBHost               string
	DBPort               string
	DBUser               string
	DBPass               string
	DBAddr               string
	DBName               string
	JWTSecret            string
	JWTExpiration        int64
	JWTRefreshExpiration int64
}

var Env = initConfig()

func initConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		errors.LogError(err)
	}

	return Config{
		APIPort: getEnv("API_PORT", "8080"),
		DBName:  getEnv("DB_NAME", ""),
		DBUser:  getEnv("DB_USER", "root"),
		DBPass:  getEnv("DB_PASS", ""),
		DBAddr: fmt.Sprintf(
			"%s:%s",
			getEnv("DB_HOST", "127.0.0.1"),
			getEnv("DB_PORT", "3306"),
		),
		JWTSecret:            getEnv("JWT_SECRET", "secret"),
		JWTExpiration:        getEnvAsInt("JWT_EXP", 3600),
		JWTRefreshExpiration: getEnvAsInt("JWT_REFRESH_EXP", 2592000),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		integer, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return integer
	}

	return fallback

}
