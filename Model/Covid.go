package Model

import "time"

type Covid struct {
	Global struct {
		NewConfirmed   int `json:"NewConfirmed"`
		TotalConfirmed int `json:"TotalConfirmed"`
		NewDeaths      int `json:"NewDeaths"`
		TotalDeaths    int `json:"TotalDeaths"`
		NewRecovered   int `json:"NewRecovered"`
		TotalRecovered int `json:"TotalRecovered"`
	} `json:"Global"`

	Countries []CountryCovid `json:"Countries"`
}

type CountryCovid struct {
	Country        string    `json:"Country"`
	CountryCode    string    `json:"CountryCode"`
	Slug           string    `json:"Slug"`
	NewConfirmed   int       `json:"NewConfirmed"`
	TotalConfirmed int       `json:"TotalConfirmed"`
	NewDeaths      int       `json:"NewDeaths"`
	TotalDeaths    int       `json:"TotalDeaths"`
	NewRecovered   int       `json:"NewRecovered"`
	TotalRecovered int       `json:"TotalRecovered"`
	Date           time.Time `json:"Date"`
}
