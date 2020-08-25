package Middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	APIKey = "SUPER_SECRET_API_KEY"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		validateToken(c)
		c.Next()
	}
}

func validateToken(c *gin.Context) {
	token := c.Request.Header.Get("X-Auth-Token")

	if token == "" {
		c.String(http.StatusOK, "API-Key required")
		c.AbortWithStatus(401)
	} else if checkToken(token) {
		c.Next()
	} else {
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.String(http.StatusOK, "Wrong API-Key... you'll never guess it, its a SUPER_SECRET API KEY!")
			c.AbortWithStatus(401)
		}
	}
}

func checkToken(token string) bool {
	if token == APIKey {
		return true
	}

	return false
}
