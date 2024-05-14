package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/HiroAcocoro/meal-planner-app-server/config"
	"github.com/HiroAcocoro/meal-planner-app-server/internal/types"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserIdByToken(token *jwt.Token) string {
	claims := token.Claims.(jwt.MapClaims)
	return claims["userId"].(string)
}

func ParseJwt(r *http.Request) (*jwt.Token, error) {
	tokenAuth := r.Header.Get("Authorization")
	if tokenAuth != "" && !strings.HasPrefix(tokenAuth, "Bearer ") {
		return nil, fmt.Errorf("invalid token")
	}

	encodedToken := strings.TrimPrefix(tokenAuth, "Bearer ")
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Env.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func CreateJwt(secret []byte, userId string) (string, error) {
	expiration := time.Second * time.Duration(config.Env.JWTExpiration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserIdFromContext(ctx context.Context) string {
	userId, ok := ctx.Value(types.UserKey).(string)
	if !ok {
		return ""
	}

	return userId
}
