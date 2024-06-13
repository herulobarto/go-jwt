package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/herulobarto/go-jwt/configs"
	"github.com/herulobarto/go-jwt/routes"
)

func main() {
	configs.ConnectDB()

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	routes.AuthRoutes(router)

	log.Println("Server running at port 7070")
	http.ListenAndServe(":7070", router)
}
