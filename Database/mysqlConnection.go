package database

import (
	"CanITravelTo/Model"
	"CanITravelTo/Utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func CreateConnection() *sql.DB {

	dbURL := Utils.ReadCredentials()

	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		fmt.Println(err.Error())
		defer db.Close()

		err = db.Ping()
		fmt.Println(err)
		if err != nil {
			fmt.Println("MySQL db is not connected")
			fmt.Println(err.Error())
		}
	}
	fmt.Println("db is connected")
	return db
}

func SelectCountry(db *sql.DB, destination string, origin string) Model.InfoCountry {
	var query strings.Builder
	query.WriteString("SELECT ")
	query.WriteString(destination)
	query.WriteString(" FROM PassportInfo WHERE Passport LIKE '")
	query.WriteString(origin)
	query.WriteString("'")

	finalQuery := query.String()

	selDB, err := db.Query(finalQuery)

	if err != nil {
		panic(err.Error())
	}
	var country Model.InfoCountry
	var info string

	for selDB.Next() {
		selDB.Scan(&info)

		if err != nil {
			panic(err.Error())
		}
		country.Info = info
	}

	return country
}
