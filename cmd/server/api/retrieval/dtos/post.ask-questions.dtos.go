package retrieval_dtos

type AskQuestionsResponse struct {
	Output string `json:"output"`
}

type AskQuestionsBody struct {
	Input string `json:"input"`
}
