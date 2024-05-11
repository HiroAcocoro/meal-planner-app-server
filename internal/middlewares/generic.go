package middlewares

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func AllowCors(next http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	log.Println("⚠️  WARNING CORS SET TO *")
	log.Println("Enabled CORS")
	return c.Handler(next)
}
