package config  // Define package config

import (
	"github.com/jinzhu/gorm"  // Import the GORM library for object-relational mapping with databases
	_ "github.com/jinzhu/gorm/dialects/mysql"  // Import MySQL dialect for GORM; using a blank identifier (_) to only run its init function
)

var (
	db *gorm.DB  // Declare a global variable db of type *gorm.DB to hold the database handle
)

// Connect function initializes a new database connection.
// Make sure to set your own database connection information.
func Connect() {
	// Open a new database connection.
	// Replace the empty string with your database connection string.
	d, err := gorm.Open("mysql", "")  // Open a new MySQL database connection
	if err != nil {  // Check for errors
		panic("failed to connect database")  // Panic and terminate if connection fails
	}
	db = d  // Set the global db variable to the new database handle
}

// GetDB function returns the database handle.
func GetDB() *gorm.DB {  // Define function GetDB that returns a pointer to a gorm.DB object
	return db  // Return the global db variable
}


