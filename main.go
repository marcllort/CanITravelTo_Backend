package main

import (
	"CanITravelTo/Controller"
	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func main() {
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /travel?destination=Spain&origin=USA
	Controller.InitHandler()
	router.GET("/travel", Controller.HandleRequest)
	router.Run(PORT)
}
