package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nealwolff/provoWorkshop/handlers"
)

func main() {

	theRouter := mux.NewRouter()

	theRouter.HandleFunc("/route", handlers.BasicHandler).Methods(http.MethodGet)
}
