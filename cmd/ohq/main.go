package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/migopp/ohq/internal/cli"
	"github.com/migopp/ohq/internal/db"
	"github.com/migopp/ohq/internal/server"
	"github.com/migopp/ohq/internal/students"
)

// System initialization
func init() {
	// Environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// DB
	db.Spawn()

	// CLI
	cli.Init()
}

// Server entry
func main() {
	// Load student list
	if cli.LSL {
		students.Load()
	}

	// Spawn the server
	if err := server.Spawn(); err != nil {
		log.Fatal("OHQ server shut down")
	}
}
