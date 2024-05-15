package middlewares

import (
	"net/http"

	"github.com/rs/cors"
)

func AllowCors(next http.Handler, h *Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	return c.Handler(next)
}
