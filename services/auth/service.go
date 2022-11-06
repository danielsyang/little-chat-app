package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(userId string) (string, error) {

	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(60)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("JWT_SECRET"))
}
