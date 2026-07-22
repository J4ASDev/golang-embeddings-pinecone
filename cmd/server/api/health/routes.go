package health

import (
	health_controllers "golang-embeddings-pinecone/cmd/server/api/health/controllers"

	"github.com/gin-gonic/gin"
)

func HealthRoutes(healthGroup *gin.RouterGroup) {
	healthGroup.GET("/ping", health_controllers.GetPingController)
}
