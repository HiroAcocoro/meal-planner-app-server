package user

import (
	"fmt"
	"net/http"

	"github.com/HiroAcocoro/meal-planner-app-server/config"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/services/auth"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/types"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /signin", h.handleSignin)
	router.HandleFunc("POST /signup", h.handleSignup)
	router.HandleFunc("POST /refresh-token", h.handleRefreshToken)
}

func (h *Handler) handleSignin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf("not found, invalid email or password"),
		)
		return
	}

	if !auth.ComparePasswords(user.Password, []byte(payload.Password)) {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf("not found, invalid email or password"),
		)
		return
	}

	// create token
	secret := []byte(config.Env.JWTSecret)
	token, err := auth.CreateJwt(secret, user.Id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// create cookie
	cookie, err := auth.CreateJwtCookie(secret, user.Id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(w, &cookie)

	utils.WriteJSON(
		w,
		http.StatusOK,
		map[string]string{"accessToken": token, "refreshToken": cookie.Value},
	)
}

func (h *Handler) handleSignup(w http.ResponseWriter, r *http.Request) {
	var payload types.SignupUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// check if user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf("user with email %s already exists", payload.Email),
		)
		return
	}

	// hash password
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// UUID
	id := uuid.New().String()

	// create user
	err = h.store.CreateUser(types.User{
		Id:       id,
		Email:    payload.Email,
		Password: hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleRefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshTokenCookie, err := r.Cookie("refreshToken")
	if err != nil {
		if err == http.ErrNoCookie {
			utils.WriteError(w, http.StatusUnauthorized, err)
			return
		}
		utils.WriteError(w, http.StatusUnauthorized, err)
    return
	}

	refreshToken := refreshTokenCookie.Value

	// Validate the refresh token
	if !auth.IsValidRefreshToken(refreshToken) {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid refresh token"))
    return
	}

	secret := []byte(config.Env.JWTSecret)

	parsedRefreshToken, err := auth.ParseRefreshJwt(refreshToken)
	
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)

	}

	userId := auth.GetUserIdByToken(parsedRefreshToken)

	// create token
	token, err := auth.CreateJwt(secret, userId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// create cookie
	cookie, err := auth.CreateJwtCookie(secret, userId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(w, &cookie)

	utils.WriteJSON(
		w,
		http.StatusOK,
		map[string]string{"accessToken": token, "refreshToken": cookie.Value},
	)
}
