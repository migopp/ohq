package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// `Spawn` starts a server at a designated address on the
// host machine. This blocks the main goroutine until the
// server fails.
func Spawn() {
	// Configure the router
	//
	// Basically, set up the controllers for each acceptable request type.
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", loginAuth, getHome)
	r.GET("/login", getLogin)
	r.POST("/login", postLogin)
	r.GET("/queue", getQueue)
	r.POST("/queue", loginAuth, postQueue)
	r.DELETE("/queue", loginAuth, adminAuth, deleteQueue)
	r.GET("/admin", loginAuth, adminAuth, getAdmin)

	// Load the templates from `web/templates/` into the engine,
	// as well as `web/components/`, and static assets
	r.Static("/static", "./web/static/")
	r.LoadHTMLGlob("web/templates/*")

	// Boot up the server
	//
	// This is just a simple `http` server for now,
	// but we _may_ want `https` support in the future.
	//
	// More info/example here:
	// https://pkg.go.dev/net/http#ListenAndServeTLS
	r.SetTrustedProxies(nil)
	sa := fmt.Sprintf("%s:%s", os.Getenv("IP"), os.Getenv("PORT"))
	fmt.Printf("Starting server at %s \n", sa)
	if err := r.Run(sa); err != nil {
		log.Fatalf("Server shut down: %v", err)
	}
}
