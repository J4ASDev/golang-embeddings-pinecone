package health_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPingController(c *gin.Context) {
	var message string = "pong"
	var status int = 200

	c.JSON(http.StatusCreated, gin.H{
		"message": message,
		"status":  status,
	})
}
