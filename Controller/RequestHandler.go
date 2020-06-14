package Controller

import (
	database "CanITravelTo/Database"
	"CanITravelTo/Model"
	"CanITravelTo/Utils"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var db *sql.DB

func InitHandler(creds string) {
	db = database.CreateConnection(creds)

}

func HandleRequest(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Saving the JSON in the corresponding fields
	var m Model.APIRequest
	err = json.Unmarshal(body, &m)
	if err != nil {
		panic(err)
	}

	destination := m.Destination
	origin := m.Origin

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
