package middleware

import (
	"context"
	"jwtauth/internal/utils"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userEmail, err := utils.ValidateJWT(r)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "userEmail", userEmail)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
