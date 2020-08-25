package Controller

import (
	database "CanITravelTo/BusinessHandler/Database"
	"CanITravelTo/BusinessHandler/Model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func HandleResponse(c *gin.Context, destination string, origin string) {

	var country Model.InfoCountry
	var countryCovid Model.CountryCovid

	code := 200
	allowed := false
	info := ""
	error := ""

	if database.ExistsCountry(db, destination) {
		// Check borders in the DB --> Return list of countries that can travel there
		country = database.SelectCountryOriginDest(db, destination, origin)
		if database.ExistsCountry(db, destination) {
			countryCovid = database.SelectCountryCovid(db, destination)
		} else {
			code = 400
			allowed = false
			error = "Destination country (covid) does NOT exist. "
		}
	} else {
		code = 400
		allowed = false
		error += "Destination country does NOT exist. "
	}

	if origin != "_" {
		if database.ExistsCountry(db, origin) {

		} else {
			code = 400
			allowed = false
			error += "Origin country does NOT exist."
		}
	}

	if country.Info != "VR" && code == 200 {
		allowed = true
		info += "You can travel to " + destination + "! There are " + fmt.Sprintf("%d", countryCovid.NewConfirmed) + " new daily COVID-19 cases. Your VISA allows you to be there for "
		if country.Info == "VF" {
			info += "an undefined amount of time!"
		} else {
			info += country.Info + " days without any special permit!"
		}
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.Header("X-Auth-Token", "SUPER_SECRET_API_KEY")
	c.JSON(code, gin.H{
		"destination": destination,
		"origin":      origin,
		"allowed":     allowed,
		"passport":    country.Info,
		"info":        info,
		"error":       error,
		"covid":       countryCovid,
	})

}
