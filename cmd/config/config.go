package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Load() *gin.Engine {
	var router *gin.Engine = gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
		},
		ExposeHeaders: []string{"Content-Length"},
	}))

	var apiGroup *gin.RouterGroup = router.Group("/api")
	var v1Group *gin.RouterGroup = apiGroup.Group("/v1")

	LoadRoutes(v1Group)

	return router
}
