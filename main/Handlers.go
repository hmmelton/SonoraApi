package main

import (
	"encoding/json"
	"net/http"
	"gopkg.in/mgo.v2/bson"
)

/*
 * Recipe Methods
 */

var recipes []Recipe

func GetRecipeEndpoint(w http.ResponseWriter, req *http.Request)  {
	// Get query parameters
	params := req.URL.Query()

	result := Recipe{}
	err := Recipes.Find(bson.M{"user_id": params.Get("id")}).One(&result)
	if err != nil {
		// There was an error -- return it to the sender, and exit
		errorMessage := map[string]string{"error": err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	// Match was found -- return the recipe
	json.NewEncoder(w).Encode(result)
}

func PostRecipeEndpoint(w http.ResponseWriter, req *http.Request) {
	recipe := Recipe{}
	// Decode recipe into variable
	err := json.NewDecoder(req.Body).Decode(&recipe)
	if err != nil {
		// There was an error
		json.NewEncoder(w).Encode(bson.M{"error": err.Error()})
		return
	}
	// Add to database
	err1 := Recipes.Insert(&recipe)

	if err1 != nil {
		// There was an error
		json.NewEncoder(w).Encode(bson.M{"error": err1.Error()})
		return
	}

	// Insert successful
	json.NewEncoder(w).Encode(bson.M{"success": "recipe added successfully"})
}

func DeleteRecipeEndpoint(w http.ResponseWriter, req *http.Request) {

}

/*
 * Menu Methods
 */

func GetMenuEndpoint(w http.ResponseWriter, req *http.Request) {

}

func PostMenuEndpoint(w http.ResponseWriter, req *http.Request) {

}

func DeleteMenuEndpoint(w http.ResponseWriter, req *http.Request) {

}