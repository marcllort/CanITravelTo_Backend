package main

import (
	"CanITravelTo/DataRetriever/Controller"
	"github.com/jasonlvhit/gocron"
	"os"
)

func main() {

	creds := os.Args[1]
	dbpass := os.Args[2]

	Controller.InitDatabase(creds, dbpass)
	Controller.CovidRetrieval()

	gocron.Every(1).Day().At("10:30:01").Do(Controller.CovidRetrieval)
	<-gocron.Start()

}
