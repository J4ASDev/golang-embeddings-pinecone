package embeddings_controllers

import (
	embeddings_dtos "golang-embeddings-pinecone/cmd/server/api/embeddings/dtos"
	embeddings_services "golang-embeddings-pinecone/cmd/server/api/embeddings/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCreateEmbeddingsController(c *gin.Context) {
	var response embeddings_dtos.CreateEmbeddingsResponseDto
	var err error

	var b embeddings_dtos.CreateEmbeddingsBodyDto

	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	response, err = embeddings_services.PostCreateEmbeddingsService(c, b.Text)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
	}

	c.JSON(http.StatusCreated, response)
}
