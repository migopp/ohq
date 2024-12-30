package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func requireAuth(c *gin.Context) {
	// Pick up `Authorization` from cookies
	toks, err := c.Cookie("Authorization")
	if err != nil {
		c.Header("hx-redirect", "/login")
		c.Status(http.StatusOK)
		return
	}

	// Parse the JWT
	token, err := jwt.Parse(toks, func(token *jwt.Token) (interface{}, error) {
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
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Expiration
		if expf, ok := claims["exp"].(float64); ok {
			expt := time.Unix(int64(expf), 0)
			if time.Now().After(expt) {
				// Need to re-login
				log.Println("JWT expired. Need to re-login.")
				c.Header("hx-redirect", "/login")
				c.Status(http.StatusOK)
				return
			}
		} else {
			// Need to re-login
			log.Println("JWT `exp` claim is missing or not valid")
			c.Header("hx-redirect", "/login")
			c.Status(http.StatusOK)
			return
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Move on along
	log.Printf("HERE! DONE!\n")
	c.Next()
}
