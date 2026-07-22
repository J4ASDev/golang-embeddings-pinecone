package config

import (
	"golang-embeddings-pinecone/cmd/server/api/embeddings"
	"golang-embeddings-pinecone/cmd/server/api/health"
	"golang-embeddings-pinecone/cmd/server/api/retrieval"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(v1Group *gin.RouterGroup) {
	var healthGroup *gin.RouterGroup = v1Group.Group("/health")
	health.HealthRoutes(healthGroup)

	var embeddingsGroup *gin.RouterGroup = v1Group.Group("/embeddings")
	embeddings.EmbeddingsRoutes(embeddingsGroup)

	var retrievalGroup *gin.RouterGroup = v1Group.Group("/retrieval")
	retrieval.RetrievalRoutes(retrievalGroup)
}
