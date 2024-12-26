package server

import (
	"fmt"
	"net/http"
)

// Server credentials
//
// Let's just hardcode the IP and port for simplicity.
var serverAddress = "0.0.0.0:6969"

// `Spawn` starts a server at a designated address on the
// host machine. If there is a failure, it returns an `error`.
func Spawn() error {
	// Configure the router
	//
	// We can just use the default router because there is
	// absolutely nothing fancy going on here.
	http.HandleFunc("GET /", getHome)
	http.HandleFunc("POST /add", postAdd)

	// Boot up the server
	//
	// This is just a simple `http` server for now,
	// but we _may_ want `https` support in the future.
	//
	// More info/example here:
	// https://pkg.go.dev/net/http#ListenAndServeTLS
	fmt.Printf("Starting server at %s \n", serverAddress)
	return http.ListenAndServe(serverAddress, nil)
}
