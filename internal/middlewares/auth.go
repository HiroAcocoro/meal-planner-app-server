package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/HiroAcocoro/meal-planner-app-server/internal/common/errors"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/services/auth"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/types"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/utils"
)

func IsAuthenticated(next http.Handler, h *Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// detect if pre-auth routes
		preAuthPaths := []string{"/signin", "/signup", "refresh-token"}
		preAuthRouter := false
		for _, path := range preAuthPaths {
			if strings.Contains(r.URL.Path, path) {
				preAuthRouter = true
				break
			}
		}

		// skip auth logic
		if preAuthRouter {
			next.ServeHTTP(w, r)
			return
		}

		token, err := auth.ParseJwt(r)
		if err != nil {
			errors.LogError(err)
			utils.WriteUnauthenticated(w)
			return
		}

		expired, err := auth.CheckTokenExpiration(token)
		if err != nil || expired {
			utils.WriteUnauthenticated(w)
			return
		}

		userIdStr := auth.GetUserIdByToken(token)

		user, err := h.userStore.GetUserById(userIdStr)
		if err != nil {
			errMsg := fmt.Errorf("failed to get user by id: %v", err)
			errors.LogError(errMsg)
			utils.WriteUnauthenticated(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, types.UserKey, user.Id)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
