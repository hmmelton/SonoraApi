package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	SetUpMongo()
	// Create router
	router := mux.NewRouter()

	// Recipe endpoints
	router.HandleFunc("/recipe", GetRecipeEndpoint).Methods("GET")
	router.HandleFunc("/recipe", PostRecipeEndpoint).Methods("POST")
	router.HandleFunc("/recipe", DeleteRecipeEndpoint).Methods("DELETE")
	router.HandleFunc("/recipe", UpdateRecipeEndpoint).Methods("PUT")
	// Menu endpoints
	router.HandleFunc("/menu", GetMenuEndpoint).Methods("GET")
	router.HandleFunc("/menu", PostMenuEndpoint).Methods("POST")
	router.HandleFunc("/menu", DeleteMenuEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":12345", router))
}