package Controller

import (
	database "CanITravelTo/Database"
	"CanITravelTo/Model"
	"CanITravelTo/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHandlerTest(c *gin.Context) {
	c.String(http.StatusOK, "Hello visitor")
}

func GetHandlerTravel(c *gin.Context) {
	c.String(http.StatusOK, "Only POST requests enabled for this endpoint. Go to canitravelto.com!")
}

func GetHandler(c *gin.Context) {

	destination := c.DefaultQuery("destination", "Spain")
	origin := c.DefaultQuery("origin", "_")

	var country Model.InfoCountry

	code := 200
	allowed := false
	info := ""

	if Utils.Has(destination) {
		info = "Destination country exists "
		// Check borders in the DB --> Return list of countries that can travel there
		country = database.SelectCountry(db, destination, origin)
	} else {
		code = 400
		allowed = false
		info = "Destination country does NOT exist "
	}

	if origin != "_" {
		if Utils.Has(origin) {
			info += "- Origin country exists"
		} else {
			code = 400
			allowed = false
			info += "- Origin country does NOT exist"
		}
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(code, gin.H{
		"destination": destination,
		"origin":      origin,
		"allowed":     allowed,
		"passport":    country.Info,
		"info":        info,
	})

}
