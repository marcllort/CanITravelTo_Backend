package main

import (
	"CanITravelTo/Controller"
	"github.com/gin-gonic/gin"
	"os"
)

const PORT = ":8080"

func main() {

	creds := os.Args[1]
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /travel?destination=Spain&origin=USA
	Controller.InitHandler(creds)
	router.POST("/travel", Controller.HandleRequest)
	router.Run(PORT)
}
