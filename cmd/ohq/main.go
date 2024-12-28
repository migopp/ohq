package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/migopp/ohq/internal/server"
)

// System initialization
func init() {
	// Environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Server entry
func main() {
	// Spawn the server
	if err := server.Spawn(); err != nil {
		log.Fatal("OHQ server shut down")
	}
}
