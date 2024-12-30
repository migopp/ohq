package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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
	// Set up an exit channel
	//
	// When I send a SIGTERM to this server,
	// I want it to exit gracefully. So I'll have to
	// set that up here.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM)
	go func() { // `deinit`
		<-sc
		log.Println("Executing `deinit`")
		db.Close()
		os.Exit(0)
	}()

	// Load student list
	if cli.LSL {
		students.Load()
	}

	// Spawn the server
	if err := server.Spawn(); err != nil {
		log.Fatal("OHQ server shut down")
	}
}
