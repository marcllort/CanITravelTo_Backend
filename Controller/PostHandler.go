package Controller

import (
	database "CanITravelTo/Database"
	"CanITravelTo/Model"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var db *sql.DB

func InitDatabase(creds string) {
	db = database.CreateConnection(creds)
}

func PostHandler(c *gin.Context) {

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

	HandleResponse(c, destination, origin)

}
