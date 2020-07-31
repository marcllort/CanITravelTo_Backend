package Controller

import (
	database "CanITravelTo/DataRetriever/Database"
	"CanITravelTo/DataRetriever/Model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CovidRetrieval() {

	response, err := http.Get("https://api.covid19api.com/summary")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		var responseObject Model.Covid
		json.Unmarshal(data, &responseObject)
		fmt.Println("Retrieving data...")
		database.UpdateCovidCountry(db, responseObject)
		fmt.Println("Covid Data updated!")
	}

}
