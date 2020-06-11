package Controller

import (
	database "CanITravelTo/Database"
	"CanITravelTo/Utils"
	"database/sql"
	"github.com/gin-gonic/gin"
)

var db *sql.DB

func InitHandler() {
	db = database.CreateConnection()
}

func HandleRequest(c *gin.Context) {

	destination := c.DefaultQuery("destination", "Spain")
	origin := c.DefaultQuery("origin", "_")
	code := 200
	allowed := false
	info := ""

	if Utils.Has(destination) {
		info = "Destination country exists "
		// Check borders in the DB --> Return list of countries that can travel there
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

	c.JSON(code, gin.H{
		"destination": destination,
		"origin":      origin,
		"allowed":     allowed,
		"info":        info,
	})
}
