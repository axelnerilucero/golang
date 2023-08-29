package main

import (
	"log"
	"mux_clase05/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	mux := mux.NewRouter()
	rutap := "/api/v1/"
	mux.HandleFunc(rutap+"ejemplo", handlers.Ejemplo_get)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
