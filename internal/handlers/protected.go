package handlers

import (
	"encoding/json"
	"net/http"
)

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value("userEmail").(string)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "welcome to the protected route",
		"user":    userEmail,
	})
}
