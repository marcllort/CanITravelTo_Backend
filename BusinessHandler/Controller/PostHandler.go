package Controller

import (
	database "CanITravelTo/BusinessHandler/Database"
	"CanITravelTo/BusinessHandler/Model"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var db *sql.DB

func InitDatabase(creds, dbpass string) {
	db = database.CreateConnection(creds, dbpass)
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
