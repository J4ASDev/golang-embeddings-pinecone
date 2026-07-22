package retrieval_controllers

import (
	retrieval_dtos "golang-embeddings-pinecone/cmd/server/api/retrieval/dtos"
	retrieval_services "golang-embeddings-pinecone/cmd/server/api/retrieval/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostAskQuestions(c *gin.Context) {
	var response retrieval_dtos.AskQuestionsResponse
	var err error

	var b retrieval_dtos.AskQuestionsBody

	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	response, err = retrieval_services.PostAskQuestions(c, b.Input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messge": "Internal server error",
		})
	}

	c.JSON(http.StatusCreated, response)
}
