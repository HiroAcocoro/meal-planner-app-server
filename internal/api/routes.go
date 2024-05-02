package api

import (
  "github.com/go-chi/render"
)

func (s *Server) routes() {
  s.router.Use(render.SetContentType(render.ContentTypeJSON))

  s.router.Get("/health", s.handleGetHealth)
}
