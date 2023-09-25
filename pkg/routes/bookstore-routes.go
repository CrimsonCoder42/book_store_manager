package routes

import (
	"github.com/gorilla/mux"  // For routing HTTP requests
	"github.com/CrimsonCoder42/go-bookstore/pkg/controllers"  // Importing the controllers package
)

// RegisterBookstoreRoutes function takes a router and registers all bookstore related routes.
var RegisterBookstoreRoutes = func(router *mux.Router) {
	// HTTP POST request to create a new book.
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	// HTTP GET request to fetch all books.
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	// HTTP GET request to fetch a book by its ID.
	// {bookId} is a variable URI segment that will be matched and its value will be captured.
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	// HTTP PUT request to update a book by its ID.
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	// HTTP DELETE request to delete a book by its ID.
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
