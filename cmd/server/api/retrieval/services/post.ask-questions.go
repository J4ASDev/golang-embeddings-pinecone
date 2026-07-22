package retrieval_services

import (
	retrieval_dtos "golang-embeddings-pinecone/cmd/server/api/retrieval/dtos"
	"golang-embeddings-pinecone/cmd/server/providers"

	"github.com/gin-gonic/gin"
)

func PostAskQuestions(c *gin.Context, input string) (retrieval_dtos.AskQuestionsResponse, error) {
	var response retrieval_dtos.AskQuestionsResponse
	var err error

	pc := providers.Container.Pinecone
	oa := providers.Container.OpenAI

	embeddings, err := oa.GenerateEmbeddings(c.Request.Context(), input)

	if err != nil {
		return response, err
	}

	var contexts []string
	queryResults, err := pc.Query(c.Request.Context(), embeddings, 1)

	for _, match := range queryResults {
		textValue, ok := match.Metadata["text"]
		if !ok {
			continue
		}

		text, ok := textValue.(string)
		if !ok || text == "" {
			continue
		}

		contexts = append(contexts, text)
	}

	if err != nil {
		return response, err
	}

	output, err := oa.GenerateResponse(c.Request.Context(), input, contexts)

	response.Output = output

	return response, err
}
