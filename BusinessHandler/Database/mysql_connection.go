package database

import (
	"CanITravelTo/BusinessHandler/Model"
	"CanITravelTo/BusinessHandler/Utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func CreateConnection(creds, dbpass string) *sql.DB {

	dbURL := Utils.ReadCredentials(creds+"/creds.json", dbpass)

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
	fmt.Println("DB is connected!")

	return db
}

func SelectCountryOriginDest(db *sql.DB, destination string, origin string) Model.InfoCountry {

	destination = strings.ReplaceAll(destination, " ", "_")
	origin = strings.ReplaceAll(origin, " ", "_")

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

func ExistsCountry(db *sql.DB, countrySelect string) bool {
	var query strings.Builder
	query.WriteString("SELECT EXISTS(SELECT * FROM PassportInfo WHERE Passport LIKE '")
	query.WriteString(countrySelect)
	query.WriteString("')")

	finalQuery := query.String()

	var exists bool
	row := db.QueryRow(finalQuery)
	err := row.Scan(&exists)

	if err != nil {
		panic(err.Error())
	}

	return exists
}

func SelectCountryCovid(db *sql.DB, country string) Model.CountryCovid {
	var query strings.Builder
	query.WriteString("SELECT Country, CountryCode, Slug, NewConfirmed, TotalConfirmed, NewDeaths, TotalDeaths, NewRecovered, TotalRecovered")
	query.WriteString(" FROM CovidInfo WHERE Country LIKE '")
	query.WriteString(country)
	query.WriteString("'")

	finalQuery := query.String()
	var covidInfo Model.CountryCovid

	db.QueryRow(finalQuery).Scan(&covidInfo.Country, &covidInfo.CountryCode, &covidInfo.Slug, &covidInfo.NewConfirmed, &covidInfo.TotalConfirmed,
		&covidInfo.NewDeaths, &covidInfo.TotalDeaths, &covidInfo.NewRecovered, &covidInfo.TotalRecovered)

	return covidInfo
}

func ExistsCountryCovid(db *sql.DB, countrySelect string) bool {
	var query strings.Builder
	query.WriteString("SELECT EXISTS(SELECT * FROM CovidInfo WHERE Country LIKE '")
	query.WriteString(countrySelect)
	query.WriteString("'")

	finalQuery := query.String()

	var exists bool
	row := db.QueryRow(finalQuery)
	err := row.Scan(&exists)

	if err != nil {
		panic(err.Error())
	}

	return exists
}

func InsertCovidCountry(db *sql.DB, covid Model.Covid) {
	var query strings.Builder

	for _, element := range covid.Countries {
		query.WriteString("INSERT INTO CovidInfo(Country, CountryCode, Slug, NewConfirmed, TotalConfirmed, NewDeaths, TotalDeaths, NewRecovered, TotalRecovered) VALUES (\"")
		query.WriteString(element.Country)
		query.WriteString("\" ,'")
		query.WriteString(element.CountryCode)
		query.WriteString("' ,'")
		query.WriteString(element.Slug)
		query.WriteString("' ,")
		query.WriteString(fmt.Sprintf("%d", element.NewConfirmed))
		query.WriteString(" ,")
		query.WriteString(fmt.Sprintf("%d", element.TotalConfirmed))
		query.WriteString(" ,")
		query.WriteString(fmt.Sprintf("%d", element.NewDeaths))
		query.WriteString(" ,")
		query.WriteString(fmt.Sprintf("%d", element.TotalDeaths))
		query.WriteString(" ,")
		query.WriteString(fmt.Sprintf("%d", element.NewRecovered))
		query.WriteString(" ,")
		query.WriteString(fmt.Sprintf("%d", element.TotalRecovered))

		query.WriteString(" )")

		finalQuery := query.String()

		insert, err := db.Query(finalQuery)
		if err != nil {
			panic(err.Error())
		}

		insert.Close()
		query.Reset()

	}

}

func UpdateCovidCountry(db *sql.DB, covid Model.Covid) {
	var query strings.Builder

	for _, element := range covid.Countries {
		query.WriteString("UPDATE CovidInfo SET ")
		query.WriteString(fmt.Sprintf("NewConfirmed=%d", element.NewConfirmed))
		query.WriteString(" ,")
		query.WriteString(fmt.Sprintf("TotalConfirmed=%d", element.TotalConfirmed))
		query.WriteString(" ,")
		query.WriteString(fmt.Sprintf("NewDeaths=%d", element.NewDeaths))
		query.WriteString(" ,")
		query.WriteString(fmt.Sprintf("TotalDeaths=%d", element.TotalDeaths))
		query.WriteString(" ,")
		query.WriteString(fmt.Sprintf("NewRecovered=%d", element.NewRecovered))
		query.WriteString(" ,")
		query.WriteString(fmt.Sprintf("TotalRecovered=%d", element.TotalRecovered))
		query.WriteString(" WHERE Country=\"")
		query.WriteString(element.Country)
		query.WriteString("\"")

		finalQuery := query.String()

		insert, err := db.Query(finalQuery)
		if err != nil {
			panic(err.Error())
		}

		insert.Close()
		query.Reset()

	}

}
