// Package declaration
package main

// Import statements for the required packages
import (
	"log"  // Logging package for logging errors
	"net/http"  // HTTP server functionalities
	"github.com/gorilla/mux"  // Gorilla Mux for HTTP routing
	_ "github.com/jinzhu/gorm/dialects/mysql"  // GORM MySQL dialect; the blank identifier is used to run its init function
	"github.com/CrimsonCoder42/go-bookstore/pkg/routes"  // Custom package for registering routes
)

// Main function, entry point of the application
func main() {
	r := mux.NewRouter()  // Create a new Gorilla Mux router
	routes.RegisterBookstoreRoutes(r)  // Register bookstore routes using the custom function from pkg/routes

	// Attach the router to the default HTTP ServeMux
	http.Handle("/", r)

	// Start the HTTP server on localhost with port 9010
	// log.Fatal() will log errors returned from ListenAndServe and terminate the program
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}


