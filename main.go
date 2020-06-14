package main

import (
	"CanITravelTo/Controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

const PORT = ":443"

func main() {

	creds := os.Args[1]
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /travel?destination=Spain&origin=USA
	Controller.InitHandler(creds)
	router.POST("/travel", Controller.HandleRequest)
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello visitor")
	})

	router.RunTLS(PORT, "Creds/https-server.crt", "Creds/https-server.key")

}
