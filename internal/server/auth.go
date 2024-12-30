package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/migopp/ohq/internal/state"
)

// `getAuth` fetches the `Authorization` cookie, used for such a purpose.
// If no such cookie exists, or it cannot be fetched, an error is returned.
func getAuth(c *gin.Context) (string, error) {
	return c.Cookie("Authorization")
}

// `getSession` parses the user's `Authorization` cookie to fetch their JWT,
// then parsing it and returning the session within.
func getSession(c *gin.Context) (state.Session, error) {
	var session state.Session
	var err error

	// Pick up `Authorization` from cookies
	toks, err := c.Cookie("Authorization")
	if err != nil {
		return session, err
	}

	// Parse the JWT
	tok, err := jwt.Parse(toks, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return session, err
	}

	// Get the claims, then parse them for useful info and pluck into session
	claims, _ := tok.Claims.(jwt.MapClaims)
	csid, cok := claims["csid"].(string)
	admin, aok := claims["admin"].(bool)
	if !cok || !aok {
		return session, fmt.Errorf("Invalid claims")
	}
	return state.Session{
		CSID:  csid,
		Admin: admin,
	}, err
}

// `loginAuth` is middleware that ensures a user is logged in (has an appropriate
// session with the server) before they gain access to a page or functionality.
func loginAuth(c *gin.Context) {
	// Pick up `Authorization` from cookies
	toks, err := c.Cookie("Authorization")
	if err != nil {
		// Plain HTTP redirection
		http.Redirect(c.Writer, c.Request, "/login", http.StatusFound)
		return
	}

	// Parse the JWT
	tok, err := jwt.Parse(toks, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Printf("JWT parsing error: %v\n", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Now we actually check the claims
	if claims, ok := tok.Claims.(jwt.MapClaims); ok {
		// Expiration
		if expf, ok := claims["exp"].(float64); ok {
			expt := time.Unix(int64(expf), 0)
			if time.Now().After(expt) {
				// Need to re-login
				log.Println("JWT expired. Need to re-login.")
				http.Redirect(c.Writer, c.Request, "/login", http.StatusFound)
				return
			}
		} else {
			// Need to re-login
			log.Println("JWT `exp` claim is missing or not valid")
			http.Redirect(c.Writer, c.Request, "/login", http.StatusFound)
			return
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Move on along
	c.Next()
}

// `adminAuth` is middleware that ensures a user is an admin before they gain access
// to a page or functionality.
func adminAuth(c *gin.Context) {
	// Get session
	se, err := getSession(c)
	if err != nil {
		http.Redirect(c.Writer, c.Request, "/login", http.StatusFound)
		return
	}

	// Ensure admin
	if !se.Admin {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Move on along
	c.Next()
}
