package models

import (
	"github.com/jinzhu/gorm"  // Importing GORM library for object-relational mapping
	"github.com/CrimsonCoder42/go-bookstore/pkg/config"  // Custom configuration package
)

var db *gorm.DB  // Database handle, global variable

// Book struct represents the model for a book with GORM annotations.
type Book struct {
	gorm.Model  // Embedding gorm.Model adds fields like ID, CreatedAt, UpdatedAt, DeletedAt
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// Initialize the package by connecting to the database and auto-migrating the Book model.
func init() {
	config.Connect()  // Connect to database
	db = config.GetDB()  // Get database handle
	db.AutoMigrate(&Book{})  // Auto-create or update table to match Book struct
}

// CreateBook method creates a new Book record in the database.
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)  // Check if it's a new record (new primary key)
	db.Create(&b)  // Insert the new book record
	return b  // Return the created book
}

// GetAllBooks fetches all book records from the database.
func GetAllBooks() []Book {
	var Books []Book  // Slice to hold the list of Books
	db.Find(&Books)  // Find and fetch all books
	return Books  // Return all books
}

// GetBookById fetches a book by its ID.
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book  // Book struct to hold the fetched book
	db := db.Where("ID = ?", Id).Find(&getBook)  // Fetch book by ID
	return &getBook, db  // Return the fetched book and db context
}

// DeleteBook deletes a book by its ID.
func DeleteBook(Id int64) Book {
	var book Book  // Book struct to hold the deleted book
	db.Where("ID = ?", Id).Delete(book)  // Delete the book by ID
	return book  // Return the deleted book (this is likely not needed)
}



