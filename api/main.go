package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	hand "github.com/nealwolff/provoWorkshop/handlers"
)

func main() {

	theRouter := mux.NewRouter()

	theRouter.HandleFunc("/route", hand.BasicHandler).Methods(http.MethodGet)
	theRouter.HandleFunc("/users", hand.UserHandler).Methods(http.MethodPost)

	log.Println("The API is listening")
	http.ListenAndServe(":8080", theRouter)
}
