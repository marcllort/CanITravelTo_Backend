package Controller

import (
	database "CanITravelTo/DataRetriever/Database"
	"CanITravelTo/DataRetriever/Model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CovidRetrieval() int {

	response, err := http.Get("https://api.covid19api.com/summary")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return -1
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		var responseObject Model.Covid
		json.Unmarshal(data, &responseObject)
		fmt.Println("Retrieving data...")
		database.UpdateCovidCountry(db, responseObject)
		fmt.Println("Covid Data updated!")
	}

	return 0
}
