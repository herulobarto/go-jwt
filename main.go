package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/herulobarto/go-jwt/configs"
)

func main() {
	configs.ConnectDB()

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	log.Println("Server running at port 8080")
	http.ListenAndServe(":8080", router)
}
