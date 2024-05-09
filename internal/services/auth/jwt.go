package auth

import (
	"time"

	"github.com/HiroAcocoro/meal-planner-app-server/config"
	"github.com/golang-jwt/jwt/v5"
)

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
