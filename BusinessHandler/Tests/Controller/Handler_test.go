package Controller

import (
	"CanITravelTo/BusinessHandler/Controller"
	"flag"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

var password string

func init() {
	flag.StringVar(&password, "password", "", "Database Password")
}

func TestHandler(t *testing.T) {

	Controller.InitDatabase("../../Creds", password)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	Controller.HandleResponse(c, "Spain", "France")

	if c.Writer.Status() != 200 {
		t.Errorf("Response was incorrect, got: %d, want: %d.", c.Writer.Status(), 200)
	}

}
