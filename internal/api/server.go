package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/config"

  "github.com/go-chi/chi/v5"
)

type Server struct {
	cfg    config.HTTPServer
	router *chi.Mux
}

func NewServer(cfg config.HTTPServer) *Server {
  srv := &Server{
    cfg: cfg,
    router: chi.NewRouter(),
  }
  srv.routes()
  return srv
}

func (s *Server) Start(ctx context.Context) {
  server := http.Server{
    Addr: fmt.Sprintf(":%d", s.cfg.Port),
    Handler: s.router,
    IdleTimeout: s.cfg.IdleTimeout,
    ReadTimeout: s.cfg.ReadTimeout,
    WriteTimeout: s.cfg.WriteTimeout,
  }

  shutdownComplete := handleShutdown(func() {
    if err := server.Shutdown(ctx); err != {
      log.Printf("server.Shutdown failed: %v\n", err)
    }
  })

  if err := server.ListenAndServe(); err == http.ErrServerClosed {
    <-shutdownComplete
  } else {
    log.Printf("http.ListenAndServe failed: %v\n", err)
  }

  log.Println("Shutdown gracefully")

}

func handleShutdown(onShutdownSignal func()) <- chan struct{} {
  shutdown := make(chat struct{})

  go func() {
    shutdownSignal := make(chan os.Signal, 1)
    signal.Notify(shutdownSignal, os.Interrupt,  syscal.SIGTERM)

    <-shutdownSignal

    onShutdownSignal()
    close(shutdown)
  }()

  return shutdown
}
