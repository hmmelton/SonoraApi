package main

import (
	"encoding/json"
	"net/http"
	"gopkg.in/mgo.v2/bson"
)

/*
 * Recipe Methods
 */

func GetRecipeEndpoint(w http.ResponseWriter, req *http.Request)  {
	// Get query parameters
	params := req.URL.Query()

	result := Recipe{}
	err := Recipes.Find(bson.M{"id": params.Get("id")}).One(&result)
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
	// Grab query parameters
	params := req.URL.Query()
	// Delete recipe
	err := Recipes.Remove(bson.M{"featured": params.Get("id")})
	if err != nil {
		// Error encountered while removing item
		json.NewEncoder(w).Encode(bson.M{"error": err.Error()})
		return
	}

	// Recipe was removed
	json.NewEncoder(w).Encode(bson.M{"success": "recipe deleted successfully"})
}

func UpdateRecipeEndpoint(w http.ResponseWriter, req *http.Request) {
	// Grab query parameters
	params := req.URL.Query()

	recipe := Recipe{}
	// Decode body of request into Recipe object
	err := json.NewDecoder(req.Body).Decode(&recipe)
	if err != nil {
		// There was an error parsing the JSON
		json.NewEncoder(w).Encode(bson.M{"decoder error": err.Error()})
		return
	}

	// Specify query as id = ID passed in query params
	query := bson.M{"id": params.Get("id")}

	// Update recipe
	err1 := Recipes.Update(query, &recipe)
	if err1 != nil {
		// There was an error updating the database
		json.NewEncoder(w).Encode(bson.M{"updating error": err1.Error()})
		return
	}

	// Update was successful -- notify user
	json.NewEncoder(w).Encode(bson.M{"success": "recipe updated successfully"})
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