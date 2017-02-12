package src

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create router
	router := mux.NewRouter()

	router.HandleFunc("/recipe", GetRecipeEndpoint).Methods("GET")
	router.HandleFunc("/recipe", PostRecipeEndpoint).Methods("POST")
	router.HandleFunc("/recipe", DeleteRecipeEndpoint).Methods("DELETE")
}