package main

import (
  "context"
  "log"

	"github.com/HiroAcocoro/meal-planner-app-server/internal/api"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/config"
)

func main() {
  ctx := context.Background()
  cfg, err := config.Load()
  if err != nil {
    log.Fatal(err)
  }

  server := api.NewServer(cfg.HTTPServer)
  server.Start(ctx)
}
