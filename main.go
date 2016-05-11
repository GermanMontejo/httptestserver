package main

import (
	"net/http"

	. "github.com/GermanMontejo/httptestserver/handlers"
	"github.com/gorilla/mux"
	"log"
)

func SetUserRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users", GetUsers).Methods("GET")
	return r
}

func main() {
	log.Println("Listening on port 8080.")
	http.ListenAndServe(":8080", SetUserRoutes())
}
