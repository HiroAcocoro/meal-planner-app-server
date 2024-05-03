# meal-planner-app-server

## Getting Started

1. In the base directory, run `cp .env.example .env` to make an env file in the root directory
2. In the base directory, run `docker-compose -f deployments/docker-compose.yml -p meal-planner up -d` to start the services in detached mode

## Running local dev

1. In the base directory, run `go mod tidy` to download necessary modules
2. run `go run cmd/meal-planner-app-server/main.go`