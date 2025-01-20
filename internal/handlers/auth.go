package handlers

import (
	"encoding/json"
	"jwtauth/internal/models"
	"jwtauth/internal/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		user.Password = string(hashedPassword)

		if err := db.Create(&user).Error; err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "user registered successfully"})
	}
}

func LoginUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		var user models.User
		if err := db.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid email or password")
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid email or password")
			return
		}

		token, err := utils.GenerateJWT(credentials.Email)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
	}
}
