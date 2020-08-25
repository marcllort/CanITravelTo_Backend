package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OptionsHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, X-Auth-Token")
	c.JSON(http.StatusOK, struct{}{})
}
