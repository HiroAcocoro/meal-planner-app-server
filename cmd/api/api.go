package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/HiroAcocoro/meal-planner-app-server/internal/middlewares"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/services/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	// @TODO subrouters

	// user handler
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(router)

	// middleware
	stack := middlewares.CreateStack(
		middlewares.AllowCors,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: stack(router),
	}

	log.Println("🚀  Server is running on port", s.addr)

	return server.ListenAndServe()
}
