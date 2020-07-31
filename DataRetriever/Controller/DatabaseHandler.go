package Controller

import (
	database "CanITravelTo/DataRetriever/Database"
	"database/sql"
)

var db *sql.DB

func InitDatabase(creds string) {
	db = database.CreateConnection(creds)
}
