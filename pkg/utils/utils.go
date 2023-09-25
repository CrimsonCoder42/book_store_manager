package utils

import (
	"encoding/json"  // For encoding and decoding JSON data
	"io"             // For I/O operations
	"net/http"       // For HTTP server functionalities
)

// ParseBody reads the body from an HTTP request and unmarshals it into the given variable x.
// The variable x should be a pointer to the variable into which the request body will be unmarshaled.
func ParseBody(r *http.Request, x interface{}) {
	// Read all content from the request body into a byte slice called 'body'.
	// The function io.ReadAll returns an error if the read operation fails.
	if body, err := io.ReadAll(r.Body); err != nil {
		// Unmarshal the JSON data from 'body' into the variable x.
		// The function json.Unmarshal returns an error if the operation fails.
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	}
}
