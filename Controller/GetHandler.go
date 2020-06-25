package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHandlerTest(c *gin.Context) {
	c.String(http.StatusOK, "Hello visitor")
}

func GetHandlerTravel(c *gin.Context) {
	c.String(http.StatusOK, "Only POST requests enabled for this endpoint. Go to canitravelto.com!")
}

func GetHandler(c *gin.Context) {

	destination := c.DefaultQuery("destination", "Spain")
	origin := c.DefaultQuery("origin", "_")

	HandleResponse(c, destination, origin)

}
