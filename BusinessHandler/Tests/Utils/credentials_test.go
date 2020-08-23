package Utils

import (
	"CanITravelTo/BusinessHandler/Utils"
	"testing"
)

func TestReadCredentials(t *testing.T) {

	creds := "../Creds/creds.json"
	dbURL := Utils.ReadCredentials(creds, "X")

	if dbURL != "testUser:X@tcp(testURL:3306)/testDB" {
		t.Errorf("Response was incorrect, got: %#v, want: %#v.", dbURL, "testUser:X@tcp(testURL:3306)/testDB")
	}

}
