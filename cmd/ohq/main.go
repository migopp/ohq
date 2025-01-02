package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/migopp/ohq/internal/cli"
	"github.com/migopp/ohq/internal/db"
	"github.com/migopp/ohq/internal/server"
)

// System initialization
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Server entry
func main() {
	cli.Init()
	if cli.LSL { // Need to load list?
		db.LoadFromFile("students.yaml")
	}

	// Spawn the db + server connections
	db.Spawn()
	server.Spawn()
}
