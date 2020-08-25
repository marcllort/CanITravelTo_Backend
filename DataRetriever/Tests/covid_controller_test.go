package Tests

import (
	"CanITravelTo/DataRetriever/Controller"
	"flag"
	"testing"
)

var password string

func init() {
	flag.StringVar(&password, "password", "", "Database Password")
}

func TestGetHandler(t *testing.T) {

	Controller.InitDatabase("../Creds", password)
	result := Controller.CovidRetrieval()

	if result != 0 {
		t.Errorf("Response was incorrect, got: %d, want: %d.", result, 0)
	}

}
