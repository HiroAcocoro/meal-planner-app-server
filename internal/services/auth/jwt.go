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

func CheckTokenExpiration(token *jwt.Token) (bool, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["expiredAt"].(float64); ok {
			//@TODO Check if token has expired
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return true, nil
			}
			return false, nil
		}
		return true, fmt.Errorf("exp claim is missing or not a float64")
	}
	return true, fmt.Errorf("invalid token")
}

func IsValidRefreshToken(refreshToken string) bool {
	// @TODO check if this token has expired
	_, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Env.JWTSecret), nil
	})

	return err == nil
}

func ParseRefreshJwt(stringToken string) (*jwt.Token, error) {
	// @TODO redundant code in ParseJWT
	parsedRefreshToken, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Env.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return parsedRefreshToken, nil
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

func CreateJwtCookie(secret []byte, userId string) (http.Cookie, error) {
	expiration := time.Second * time.Duration(config.Env.JWTRefreshExpiration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return http.Cookie{}, err
	}

	cookie := http.Cookie{
		Name:     "refreshToken",
		Value:    tokenString,
		Expires:  time.Now().Add(expiration),
		HttpOnly: true,
		Secure:   true,
	}

	return cookie, nil
}

func GetUserIdFromContext(ctx context.Context) string {
	userId, ok := ctx.Value(types.UserKey).(string)
	if !ok {
		return ""
	}

	return userId
}
