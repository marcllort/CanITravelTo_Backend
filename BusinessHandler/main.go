package main

import (
	"CanITravelTo/BusinessHandler/Controller"
	"CanITravelTo/BusinessHandler/Middleware"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

const PORT = ":443"

func main() {

	creds := os.Args[1]
	dbpass := os.Args[2]
	myfile := os.NewFile(0, "server.log")

	gin.DefaultWriter = io.MultiWriter(myfile, os.Stdout)
	router := gin.Default()

	router.Use(Middleware.AuthMiddleware())

	Controller.InitDatabase(creds, dbpass)

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /travel?destination=Spain&origin=USA

	router.OPTIONS("/travel", Controller.OptionsHandler)
	router.POST("/travel", Controller.PostHandler)
	router.GET("/travel", Controller.GetHandler)
	router.GET("/test", Controller.GetHandlerTest)

	err := router.RunTLS(PORT, creds+"/https-server.crt", creds+"/https-server.key")

	if err != nil {
		panic(err.Error())
	}

}
