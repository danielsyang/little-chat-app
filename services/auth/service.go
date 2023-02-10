package auth

import (
	"errors"
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

func ValidateToken(tokenParam string) (string, error) {

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenParam, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("JWT_SECRET"), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return claims["userId"].(string), nil
	}

	return "", errors.New("invalid token")
}
