package Utils

import (
	"encoding/json"
	"fmt"
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

func ReadCredentials(credentials, dbpass string) string {
	file, err := ioutil.ReadFile(credentials)
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

	var creds Credentials
	creds.user = m["user"].(string)
	creds.password = dbpass
	//creds.password = m["password"].(string)
	creds.hostname = m["hostname"].(string)
	creds.port = m["port"].(string)
	creds.database = m["database"].(string)
	var url strings.Builder
	url.WriteString(creds.user)
	url.WriteString(":")
	url.WriteString(creds.password)
	url.WriteString("@tcp(")
	url.WriteString(creds.hostname)
	url.WriteString(":")
	url.WriteString(creds.port)
	url.WriteString(")/")
	url.WriteString(creds.database)

	dbURL := url.String()

	return dbURL
}
