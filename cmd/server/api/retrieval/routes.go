package retrieval

import (
	retrieval_controllers "golang-embeddings-pinecone/cmd/server/api/retrieval/controllers"

	"github.com/gin-gonic/gin"
)

func RetrievalRoutes(retrievalGroup *gin.RouterGroup) {
	retrievalGroup.POST("/ask", retrieval_controllers.PostAskQuestions)
}
