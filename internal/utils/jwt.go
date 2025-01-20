package utils

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("sss")

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	const bearerPrefix = "Bearer "
	if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		return "", errors.New("invalid authorization header format")
	}
	tokenString := authHeader[len(bearerPrefix):]

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}
	return claims.Email, nil
}
