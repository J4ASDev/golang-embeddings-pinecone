package embeddings_dtos

type CreateEmbeddingsBodyDto struct {
	Text string `json:"text" binding:"required,min=1"`
}

type CreateEmbeddingsResponseDto struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}
