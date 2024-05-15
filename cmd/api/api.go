package api

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/HiroAcocoro/meal-planner-app-server/internal/middlewares"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/services/test"
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

func (s *APIServer) Start(ctx context.Context) error {
	router := http.NewServeMux()
	// @TODO subrouters

	apiRouter := http.NewServeMux()
	apiRouter.Handle("/api/", http.StripPrefix("/api", router))

	// user handler
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(router)

	// test handler
	testHandler := test.NewHandler(userStore)
	testHandler.RegisterRoutes(router)

	// middlewares
	middlewareHandler := middlewares.NewHandler(userStore)
	middlewareStack := middlewares.CreateStack(
		middlewares.AllowCors,
		middlewares.IsAuthenticated,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareStack(apiRouter, middlewareHandler),
	}

	log.Println("🚀  Server is running on port", s.addr)

	shutdownComplete := handleShutdown(func() {
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("server.Shutdown failed: %v\n", err)
		}
	})

	if err := server.ListenAndServe(); err == http.ErrServerClosed {
		<-shutdownComplete
	} else {
		log.Printf("http.ListenAndServe failed: %v\n", err)
	}

	log.Println("😴  Shutdown gracefully")

	return server.ListenAndServe()
}

func handleShutdown(onShutdownSignal func()) <-chan struct{} {
	shutdown := make(chan struct{})

	go func() {
		shutdownSignal := make(chan os.Signal, 1)
		signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

		<-shutdownSignal

		onShutdownSignal()
		close(shutdown)
	}()

	return shutdown
}
