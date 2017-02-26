package main

import (
	"gopkg.in/mgo.v2"
	"time"
	"log"
)

// Pointer to MongoDB Session used by other files
var MongoSession *mgo.Session
// Collections
var Recipes *mgo.Collection
var Menus *mgo.Collection

// This function sets up a connection with the MongoDB
func SetUpMongo() {
	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	MongoSession = session

	// Reads may not be entirely up-to-date, but they will always see the
	// history of changes moving forward, the data read will be consistent
	// across sequential queries in the same session, and modifications made
	// within the session will be observed in following queries (read-your-writes).
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	MongoSession.SetMode(mgo.Monotonic, true)

	// Set collection variables
	Recipes = session.DB("sonora").C("recipes")
	Menus = session.DB("sonora").C("menus")
}