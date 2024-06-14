package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// ParseBody reads the request body and unmarshals it into the provided interface.
// It returns an error if any step fails.
func ParseBody(r *http.Request, x interface{}) error {
	// Ensure the body is closed after reading
	defer r.Body.Close()

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return err
	}

	// Unmarshal the JSON data into the provided interface
	if err := json.Unmarshal(body, x); err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return err
	}

	return nil
}
