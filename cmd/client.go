package main

import (
	"log"
	"syncway/internal/adapters"
)

func main() {
	// Initialize the SQLite adapter
	_, err := adapters.NewSQLiteAdapter("file:data.db")
	if err != nil {
		log.Fatalf("Failed to initialize SQLite adapter: %v", err)
	}

	//// Example: Using the adapter to track changes
	//changes, err := adapters
	//if err != nil {
	//	log.Fatalf("Failed to track changes: %v", err)
	//}
	//fmt.Printf("Tracked changes: %+v\n", changes)
}
