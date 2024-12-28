package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/migopp/ohq/internal/state"
	"github.com/migopp/ohq/internal/users"
)

// `getHome` serves a request to fetch the home page.
func getHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Users": state.GlobalState.Queue,
	})
}

// `getQueue` seerves a request to view the queue.
func getQueue(c *gin.Context) {
	c.HTML(http.StatusOK, "queue_content.html", gin.H{
		"Users": state.GlobalState.Queue,
	})
}

// `postQueue` serves a request to add a student to the queue.
func postQueue(c *gin.Context) {
	// Extract user info and offer it to the queue
	id := c.Request.PostFormValue("qid")
	u := users.User{
		ID: id,
	}
	state.GlobalState.Offer(u)

	// Serve the updates to the home page
	//
	// In this case it's to the `#queue-disp` element ID.
	// We have a hook that leads here in the HTMX, and
	// so whatever we write here replaces the current
	// contents in the DOM.
	c.HTML(http.StatusOK, "queue_content.html", gin.H{
		"Users": state.GlobalState.Queue,
	})
}

// `deleteQueue` seerves a request to poll from the queue.
func deleteQueue(c *gin.Context) {
	// Poll from the queue (if possible)
	_, err := state.GlobalState.Poll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "err.html", gin.H{
			"Err": err,
		})
	}

	// Serve the updates
	c.HTML(http.StatusOK, "queue_content.html", gin.H{
		"Users": state.GlobalState.Queue,
	})
}
