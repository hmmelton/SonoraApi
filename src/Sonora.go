package main

import (
	//"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"encoding/json"
)

type Ingredient struct {
	Name		string	`json:"name"`
	Count		float32	`json:"count"`
	CountType	string	`json:"count_type"`
}

type Recipe struct {
	Id				string			`json:"id"`
	UserId			string			`json:"user_id"`
	FirstName		string			`json:"first_name"`
	LastName		string			`json:"last_name"`
	Name			string			`json:"name"`
	Image			string			`json:"image"`
	Instructions	[]string		`json:"instructions"`
	Ingredients		[]Ingredient	`json:"ingredients"`
	Tags			[]string		`json:"tags"`
}

type Menu struct {
	Id			string		`json:"id"`
	UserId		string		`json:"user_id"`
	Image		string		`json:"image"`
	ProfImage	string		`json:"prof_image"`
	Name		string		`json:"name"`
	Appetizers	*[]Recipe	`json:"appetizers"`
	Entrees		*[]Recipe	`json:"entrees"`
	Sides		*[]Recipe	`json:"sides"`
	Desserts	*[]Recipe	`json:"desserts"`
	Beverages	*[]Recipe	`json:"beverages"`
	Tags		[]string	`json:"tags"`
}

// TODO: delete this
var recipes []Recipe

func main() {
	// Create router
	router := mux.NewRouter()

	// Recipe endpoints
	router.HandleFunc("/recipe", GetRecipeEndpoint).Methods("GET")
	/*router.HandleFunc("/recipe", PostRecipeEndpoint).Methods("POST")
	router.HandleFunc("/recipe", DeleteRecipeEndpoint).Methods("DELETE")
	// Menu endpoints
	router.HandleFunc("/menu", GetMenuEndpoint).Methods("GET")
	router.HandleFunc("/menu", PostMenuEndpoint).Methods("POST")
	router.HandleFunc("/menu", DeleteMenuEndpoint).Methods("DELETE")*/

	// TODO: delete this
	recipes = append(recipes, Recipe{Id: "1", UserId: "1", FirstName: "Harrison", LastName: "Melton",
		Name: "Fancy ABJ", Image: "/image.jpg", Instructions: []string{"Spread almond butter", "spread sliced fruit"},
		Tags: []string{"vegan", "gluten-free"}})
	recipes = append(recipes, Recipe{Id: "2", UserId: "2", FirstName: "Harrison", LastName: "Melton",
		Name: "Count Chocula", Image: "/image.jpg", Instructions: []string{"Pour Count Chocula", "pour almond milk"},
		Tags: []string{"vegan", "gluten-free"}})

	log.Fatal(http.ListenAndServe(":12345", router))
}

/*
 * Recipe Methods
 */

func GetRecipeEndpoint(w http.ResponseWriter, req *http.Request)  {
	// Get query parameters
	params := req.URL.Query()
	for _, item := range recipes {
		// Check each recipe to see if id matches parameter
		if item.Id == params.Get("id") {
			// If it matches, return the recipe
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// No match was found, return empty json
	noneFound := map[string]string{"error": "no matching entry found"}
	json.NewEncoder(w).Encode(noneFound)
}