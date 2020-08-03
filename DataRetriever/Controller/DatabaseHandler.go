package Controller

import (
	database "CanITravelTo/DataRetriever/Database"
	"database/sql"
)

var db *sql.DB

func InitDatabase(creds, dbpass string) {
	db = database.CreateConnection(creds, dbpass)
}
