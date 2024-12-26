package main

import (
	"log"

	"github.com/migopp/ohq/internal/server"
)

// Entry point
func main() {
	// Spawn the server
	if err := server.Spawn(); err != nil {
		log.Fatal("OHQ server shut down")
	}
}
