package server

import (
	"fmt"
	"net/http"

	"github.com/migopp/ohq/internal/state"
	"github.com/migopp/ohq/internal/templates"
	"github.com/migopp/ohq/internal/users"
)

// `getHome` serves a request to fetch the home page.
func getHome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving `GET /`\n")
	qc := templates.QueueContent{
		Users: state.GlobalState.Queue,
	}
	templates.ServeTemplate(templates.Home, w, r, qc)
}

// `postAdd` serves a request to add a student to the queue.
func postAdd(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving `POST /add`\n")

	// Extract user info and offer it to the queue
	err := r.ParseForm()
	if err != nil {
		es := fmt.Sprintf("Error parsing ID input [%v]", err)
		http.Error(w, es, http.StatusInternalServerError)
		fmt.Printf("%s", es)
		return
	}
	id := r.FormValue("qid")
	u := users.User{
		ID: id,
	}
	state.GlobalState.Offer(u)

	// Serve the updates to the home page
	//
	// In this case it's to the `#queue-disp` element ID.
	// We have a hook that leads here in the HTMX, and
	// so whatever we write to `w` replaces the current
	// contents in the DOM.
	qc := templates.QueueContent{
		Users: state.GlobalState.Queue,
	}
	templates.ServeTemplate(templates.QueueDisplay, w, r, qc)
}
