package config

import (
	"time"

	"github.com/joho/godotenv"
)

type Configuration struct {
	HTTPServer
}

type HTTPServer struct {
	IdleTimeout  time.Duration `godotenv:"HTTP_SERVER_IDLE_TIMEOUT"  default:"60s"`
	Port         int           `godotenv:"PORT"                      default:"8080"`
	ReadTimeout  time.Duration `godotenv:"HTTP_SERVER_READ_TIMEOUT"  default:"1s"`
	WriteTimeout time.Duration `godotenv:"HTTP_SERVER_WRITE_TIMEOUT" default:"2s"`
}

func Load() (Configuration, error) {
	cfg := Configuration
	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
