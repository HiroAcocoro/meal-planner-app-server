# Build the application from source
FROM golang:1.22.3 AS build-stage
  WORKDIR /server

  COPY go.mod go.sum ./
  RUN go mod download
  
  COPY . .
  
  RUN CGO_ENABLED=0 GOOS=linux go build -o /app-server ./cmd/server/main.go
  
# Deploy the application binary into a lean image
FROM scratch AS build-release-stage
  ARG API_PORT

  WORKDIR /
  
  COPY --from=build-stage /app-server /app-server

  EXPOSE ${API_PORT}

  ENTRYPOINT ["/app-server"]
