package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nealwolff/provoWorkshop/handlers"
)

func main() {

	theRouter := mux.NewRouter()

	theRouter.HandleFunc("/route", handlers.BasicHandler).Methods(http.MethodGet)

	log.Println("The API is listening")
	http.ListenAndServe(":8080", theRouter)
}
