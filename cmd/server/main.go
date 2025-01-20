package main

import (
	"fmt"
	"log"
	"net/http"

	"jwtauth/internal/handlers"
	"jwtauth/internal/middleware"
	"jwtauth/pkg/database"

	"github.com/gorilla/mux"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Маршруты
	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.RegisterUser(db)).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser(db)).Methods("POST")

	// Защищённые маршруты
	protected := r.PathPrefix("/protected").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/endpoint", handlers.ProtectedEndpoint).Methods("GET")

	// Запуск сервера
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
