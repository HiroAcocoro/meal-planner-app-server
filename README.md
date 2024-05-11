# meal-planner-app-server

## Getting Started

1. In the base directory, run: `cp .env.example .env` and `cp .env.example ./deployments/.env`. This will create env files necessary for running locally as well as running docker-compose


## Creating Docker Container 🐋

1. run: `make docker-compose`

This will run **mysql** and **go server** inside a docker container. Make sure to run migrations afterwards to update your database.

## Running migrations

1. run: `make migrate-up`


## Running local dev

1. run: `go mod download`
2. run: `make build`
3. run: `make run`