package main

import (
	"CanITravelTo/Controller"
	"github.com/gin-gonic/gin"
	"os"
)

const PORT = ":443"

func main() {

	creds := os.Args[1]
	router := gin.Default()
	whitelist := make(map[string]bool)
	whitelist["localhost"] = true
	whitelist["canitravelto.com"] = true

	//router.Use(Middleware.IPWhiteList(whitelist))

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /travel?destination=Spain&origin=USA
	Controller.InitHandler(creds)
	router.OPTIONS("/travel", Controller.OptionsHandler)
	router.POST("/travel", Controller.PostHandler)
	router.GET("/test", Controller.GetHandlerTest)
	router.GET("/travel", Controller.GetHandlerTravel)

	router.RunTLS(PORT, creds+"/https-server.crt", creds+"/https-server.key")

	// AutoTLS config
	/*m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("canitravelto.com", "62.57.154.24", "localhost"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(router, &m))*/

}
