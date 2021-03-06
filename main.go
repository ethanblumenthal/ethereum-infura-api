package main

import (
	"infura-challenge/app"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	projectId := os.Getenv("PROJECT_ID")
	if projectId == "" {
		log.Fatal("$PROJECT_ID must be set")
	}

	// Create routes
	router := app.NewRouter()

	// These two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"}) 
 	allowedMethods := handlers.AllowedMethods([]string{"GET"})

	// Launch server with CORS validations
 	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}