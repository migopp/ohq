package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// The actual server engine
var engine gin.Engine

// `Spawn` starts a server at a designated address on the
// host machine. If there is a failure, it returns an `error`.
func Spawn() error {
	// Configure the router
	//
	// This includes locating the static assets as well as setting
	// up the controllers for each acceptable request type.
	r := gin.Default()
	r.Static("/static", "./web/static/")
	r.GET("/", getHome)
	r.GET("/login", getLogin)
	r.POST("/login", postLogin)
	r.GET("/queue", getQueue)
	r.POST("/queue", postQueue)
	r.DELETE("/queue", deleteQueue)

	// Load the templates from `web/templates/` into the engine
	r.LoadHTMLGlob("web/templates/*")

	// Boot up the server
	//
	// This is just a simple `http` server for now,
	// but we _may_ want `https` support in the future.
	//
	// More info/example here:
	// https://pkg.go.dev/net/http#ListenAndServeTLS
	sa := fmt.Sprintf("%s:%s", os.Getenv("IP"), os.Getenv("PORT"))
	fmt.Printf("Starting server at %s \n", sa)
	return r.Run(sa)
}
