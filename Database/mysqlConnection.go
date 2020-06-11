package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
	"strings"
)

type Credentials struct {
	user     string
	password string
	hostname string
	port     string
	database string
}

func CreateConnection() *sql.DB {

	dbURL := readCredentials()

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

func readCredentials() string {
	file, err := ioutil.ReadFile("creds.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed read file: %s\n", err)
		os.Exit(1)
	}

	var f interface{}
	err = json.Unmarshal(file, &f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse JSON: %s\n", err)
		os.Exit(1)
	}

	// Type-cast `f` to a map by means of type assertion.
	m := f.(map[string]interface{})
	fmt.Printf("Parsed data: %v\n", m)

	var creds Credentials
	creds.user = m["user"].(string)
	creds.password = m["password"].(string)
	creds.hostname = m["hostname"].(string)
	creds.port = m["port"].(string)
	creds.database = m["database"].(string)
	var url strings.Builder
	url.WriteString(creds.user)
	url.WriteString(":")
	url.WriteString(creds.password)
	url.WriteString("@tcp(")
	url.WriteString(creds.hostname)
	url.WriteString("/")
	url.WriteString(creds.port)
	url.WriteString(")/")
	url.WriteString(creds.database)

	dbURL := url.String()

	return dbURL
}
