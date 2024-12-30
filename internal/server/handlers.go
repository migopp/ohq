package server

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/migopp/ohq/internal/db"
	"github.com/migopp/ohq/internal/state"
)

// `getHome` serves a request to fetch the home page.
func getHome(c *gin.Context) {
	c.HTML(http.StatusOK, "main.tmpl.html", gin.H{
		"Component": "home",
		"Users":     state.GlobalState.Queue,
	})
}

// `getLogin` serves a request to view the login page.
func getLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "main.tmpl.html", gin.H{
		"Component": "login",
	})
}

// `postLogin` serves a request to login to the `ohq` system.
func postLogin(c *gin.Context) {
	// Extract login details
	un := c.PostForm("username")
	pw := c.PostForm("password")

	// Fetch from DB and verify credentials
	u, err := db.FetchUserWithName(un)
	if err != nil {
		c.HTML(http.StatusOK, "main.tmpl.html", gin.H{
			"Component": "err",
			"Err":       err,
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
	if err != nil {
		c.HTML(http.StatusOK, "main.tmpl.html", gin.H{
			"Component": "err",
			"Err":       err,
		})
		return
	}

	// Generate and attach JWT
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u.Username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"onq": false,
	})
	toks, err := tok.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.HTML(http.StatusOK, "main.tmpl.html", gin.H{
			"Component": "err",
			"Err":       err,
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", toks, 3600*24, "", "", false, true)

	// Redirect to `/`
	c.Header("hx-redirect", "/")
	c.Status(http.StatusOK)
	return
}

// `getQueue` seerves a request to view the queue.
func getQueue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Students": state.GlobalState.Queue,
	})
}

// `postQueue` serves a request to add a student to the queue.
func postQueue(c *gin.Context) {
	// Extract user info and offer it to the queue
	claims, err := getClaims(c)
	if err != nil {
		// Unable to fetch the claims -- likely a bad JWT,
		// so re-login may fix it.
		c.Header("hx-redirect", "/login")
		c.Status(http.StatusOK)
		return
	}
	csid, ok := claims["sub"].(string)
	if !ok {
		// Again, bad JWT. Try re-login.
		c.Header("hx-redirect", "/login")
		c.Status(http.StatusOK)
		return
	}
	e := state.Entry{
		CSID: csid,
	}
	state.GlobalState.Offer(e)

	// Serve the updates to the home page
	//
	// In this case it's to the `#queue-disp` element ID.
	// We have a hook that leads here in the HTMX, and
	// so whatever we write here replaces the current
	// contents in the DOM.
	c.HTML(http.StatusOK, "components/qc", gin.H{
		"Users": state.GlobalState.Queue,
	})
}

// `deleteQueue` seerves a request to poll from the queue.
func deleteQueue(c *gin.Context) {
	// Poll from the queue (if possible)
	_, err := state.GlobalState.Poll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "main.tmpl.html", gin.H{
			"Component": "err",
			"Err":       err,
		})
		return
	}

	// Serve the updates
	c.HTML(http.StatusOK, "main.tmpl.html", gin.H{
		"Component": "qc",
		"Users":     state.GlobalState.Queue,
	})
}
