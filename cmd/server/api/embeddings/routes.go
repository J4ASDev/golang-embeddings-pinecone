package embeddings

import (
	embeddings_controllers "golang-embeddings-pinecone/cmd/server/api/embeddings/controllers"

	"github.com/gin-gonic/gin"
)

func EmbeddingsRoutes(embeddingsGroup *gin.RouterGroup) {
	embeddingsGroup.POST("/create", embeddings_controllers.PostCreateEmbeddingsController)
}
