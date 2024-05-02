package config

import (
  "os"
  "fmt"
	"time"

  "github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const envPrefix = "API"

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
  errEnv := godotenv.Load("../../.env")
  if errEnv != nil {
    fmt.Printf("Error loading .env file: %v\n", errEnv)
    os.Exit(1)
  }

	var cfg Configuration
	errCfg := envconfig.Process(envPrefix, &cfg)
	if errCfg != nil {
		return cfg, errCfg
	}

	return cfg, nil
}
