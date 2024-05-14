package middlewares

import (
	"net/http"

	"github.com/HiroAcocoro/meal-planner-app-server/internal/types"
)

type Handler struct {
	userStore types.UserStore
}

func NewHandler(
	userStore types.UserStore,
) *Handler {
	return &Handler{
		userStore:  userStore,
	}
}

type Middleware func(http.Handler, *Handler) http.Handler

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler, h *Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next, h)
		}

		return next
	}
}
