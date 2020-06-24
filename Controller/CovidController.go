package Controller

import (
	database "CanITravelTo/Database"
	"CanITravelTo/Model"
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

		database.UpdateCovidCountry(db, responseObject)
	}

}
