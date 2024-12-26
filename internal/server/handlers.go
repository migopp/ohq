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

// `getQueue` seerves a request to view the queue.
func getQueue(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving `GET /queue`\n")
	qc := templates.QueueContent{
		Users: state.GlobalState.Queue,
	}
	templates.ServeTemplate(templates.QueueDisplay, w, r, qc)
}

// `postQueue` serves a request to add a student to the queue.
func postQueue(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving `POST /queue`\n")

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

// `deleteQueue` seerves a request to poll from the queue.
func deleteQueue(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving `DELETE /queue`\n")

	// Poll from the queue (if possible)
	_, err := state.GlobalState.Poll()
	if err != nil {
		// Failure -> throw the error to the frontend
		ec := templates.ErrContent{
			Err: err,
		}
		templates.ServeTemplate(templates.Err, w, r, ec)
	}

	// Serve the updates
	qc := templates.QueueContent{
		Users: state.GlobalState.Queue,
	}
	templates.ServeTemplate(templates.QueueDisplay, w, r, qc)

}
