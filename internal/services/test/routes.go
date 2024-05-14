package test

import (
	"net/http"

	"github.com/HiroAcocoro/meal-planner-app-server/internal/services/auth"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/types"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/utils"
)

type Handler struct {
	userStore types.UserStore
}

func NewHandler(userStore types.UserStore) *Handler {
	return &Handler{
		userStore: userStore,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /test", h.handleTestRoute)
}

func (h *Handler) handleTestRoute(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r.Context())
	user, _ := h.userStore.GetUserById(userId)

	utils.WriteJSON(w, http.StatusOK, user)
}
