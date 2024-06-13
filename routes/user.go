package routes

import (
	"github.com/gorilla/mux"
	"github.com/herulobarto/go-jwt/controllers"
	"github.com/herulobarto/go-jwt/middleware"
)

func UserRoutes(r *mux.Router) {
	router := r.PathPrefix("/users").Subrouter()

	router.Use(middleware.Auth)

	router.HandleFunc("/me", controllers.Me).Methods("GET")
}
