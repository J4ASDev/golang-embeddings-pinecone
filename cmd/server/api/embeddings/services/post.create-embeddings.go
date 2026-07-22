package embeddings_services

import (
	embeddings_dtos "golang-embeddings-pinecone/cmd/server/api/embeddings/dtos"
	"golang-embeddings-pinecone/cmd/server/providers"
	pinecone_provider "golang-embeddings-pinecone/cmd/server/providers/pinecone"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostCreateEmbeddingsService(c *gin.Context, text string) (embeddings_dtos.CreateEmbeddingsResponseDto, error) {
	var response embeddings_dtos.CreateEmbeddingsResponseDto
	var err error

	pc := providers.Container.Pinecone
	oa := providers.Container.OpenAI

	var id uuid.UUID = uuid.New()

	embeddings, err := oa.GenerateEmbeddings(c.Request.Context(), text)

	var vectors = []pinecone_provider.Vector{
		{
			ID:       id.String(),
			Values:   embeddings,
			Metadata: map[string]any{"text": text},
		},
	}

	err = pc.Upsert(c.Request.Context(), vectors)

	if err != nil {
		return response, err
	}

	response.Id = id.String()
	response.Message = "Embedding stored successfully"

	return response, err
}
