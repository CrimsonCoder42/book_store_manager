// Package controllers is responsible for handling HTTP routes and operations
package controllers

// Import necessary packages
import (
	"encoding/json"  // To handle JSON data
	"fmt"  // For debugging and printing logs
	"net/http"  // For HTTP request and response handling
	"strconv"  // For converting strings to integers

	"github.com/CrimsonCoder42/go-bookstore/pkg/models"  // Custom Book model
	"github.com/CrimsonCoder42/go-bookstore/pkg/utils"  // Custom utilities
	"github.com/gorilla/mux"  // For routing HTTP requests
)

// Declare a new Book model variable
var newBook models.Book

// GetBook handles GET requests to fetch all books
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()  // Fetch all books from the model
	res, _ := json.Marshal(newBooks)  // Marshal the book data to JSON
	w.Header().Set("Content-Type", "application/json")  // Set response type to JSON
	w.WriteHeader(http.StatusOK)  // Set HTTP status to 200 OK
	w.Write(res)  // Write JSON response to the client
}

// GetBookById handles GET requests to fetch a book by ID
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)  // Get path variables from the request URL
	bookId := vars["bookId"]  // Extract book ID from path variables
	ID, err := strconv.ParseInt(bookId, 0, 0)  // Convert bookId to int64
	if err != nil {
		fmt.Println("Error while parsing")  // Log parsing error
	}
	bookDetails, _ := models.GetBookById(ID)  // Fetch book details by ID
	res, _ := json.Marshal(bookDetails)  // Marshal book details to JSON
	w.Header().Set("Content-Type", "application/json")  // Set response type to JSON
	w.WriteHeader(http.StatusOK)  // Set HTTP status to 200 OK
	w.Write(res)  // Write JSON response to the client
}

// CreateBook handles POST requests to create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}  // Declare a pointer to a new Book model
	utils.ParseBody(r, CreateBook)  // Parse the HTTP request body and fill the model
	b := CreateBook.CreateBook()  // Create a new book using the model method
	res, _ := json.Marshal(b)  // Marshal the newly created book to JSON
	w.Header().Set("Content-Type", "application/json")  // Set response type to JSON
	w.WriteHeader(http.StatusOK)  // Set HTTP status to 200 OK
	w.Write(res)  // Write JSON response to the client
}

// DeleteBook handles DELETE requests to delete a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)  // Get path variables from the request URL
	bookId := vars["bookId"]  // Extract book ID from path variables
	ID, err := strconv.ParseInt(bookId, 0, 0)  // Convert bookId to int64
	if err != nil {
		fmt.Println("Error while parsing")  // Log parsing error
	}
	book := models.DeleteBook(ID)  // Delete the book by ID using the model method
	res, _ := json.Marshal(book)  // Marshal the deleted book details to JSON
	w.Header().Set("Content-Type", "application/json")  // Set response type to JSON
	w.WriteHeader(http.StatusOK)  // Set HTTP status to 200 OK
	w.Write(res)  // Write JSON response to the client
}

// UpdateBook handles PUT requests to update a book by ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}  // Declare a pointer to a new Book model for updates
	utils.ParseBody(r, updateBook)  // Parse the HTTP request body and fill the update model
	vars := mux.Vars(r)  // Get path variables from the request URL
	bookId := vars["bookId"]  // Extract book ID from path variables
	ID, err := strconv.ParseInt(bookId, 0, 0)  // Convert bookId to int64
	if err != nil {
		fmt.Println("Error while parsing")  // Log parsing error
	}
	bookDetails, db := models.GetBookById(ID)  // Fetch existing book details by ID
	// Update book details based on non-empty fields in the request body
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)  // Save updated book details
	res, _ := json.Marshal(bookDetails)  // Marshal updated book details to JSON
	w.Header().Set("Content-Type", "application/json")  // Set response type to JSON
	w.WriteHeader(http.StatusOK)  // Set HTTP status to 200 OK
	w.Write(res)  // Write JSON response to the client
}
