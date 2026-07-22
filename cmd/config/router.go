package config

import (
	"golang-embeddings-pinecone/cmd/server/api/health"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(v1Group *gin.RouterGroup) {
	var healthGroup *gin.RouterGroup = v1Group.Group("/health")
	health.HealthRoutes(healthGroup)
}
