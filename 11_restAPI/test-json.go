package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Define the struct for testing
type Event struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"dateTime"` // Ensure proper type for ISO 8601 parsing
}

func testJson() {
	// JSON input (replace with your example)
	jsonData := `{
		"name": "Test Event",
		"description": "Test Description",
		"location": "Test Location",
		"dateTime": "2025-01-01T15:30:00.000Z"
	}`

	// Variable to hold unmarshaled data
	var event Event

	// Unmarshal JSON into the struct
	err := json.Unmarshal([]byte(jsonData), &event)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return
	}

	// Output the struct
	fmt.Printf("Unmarshaled Struct: %+v\n", event)
}
