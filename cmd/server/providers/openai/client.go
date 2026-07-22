package openai_provider

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/packages/param"
	"github.com/openai/openai-go/v3/responses"
)

type Client struct {
	oa openai.Client
}

func (c *Client) GenerateEmbeddings(ctx context.Context, text string) ([]float32, error) {
	var response []float32
	var err error

	dimension, err := strconv.ParseInt(os.Getenv("PINECONE_DIMENSION"), 10, 32)

	if err != nil {
		return response, fmt.Errorf("parse PINECONE_DIMENSION: %w", err)
	}

	result, err := c.oa.Embeddings.New(
		ctx,
		openai.EmbeddingNewParams{
			Model:          "text-embedding-3-small",
			Dimensions:     param.NewOpt(dimension),
			EncodingFormat: "float",
			Input: openai.EmbeddingNewParamsInputUnion{
				OfString: openai.String(text),
			},
		},
	)

	values64 := result.Data[0].Embedding
	values32 := make([]float32, len(values64))

	for i, v := range values64 {
		values32[i] = float32(v)
	}

	response = values32

	return response, err
}

func (c *Client) GenerateResponse(ctx context.Context, input string, contexts []string) (string, error) {
	var response string
	var err error

	var prompt string

	if len(contexts) >= 1 {
		prompt = fmt.Sprintf("Context:\n%s\n\nQuestion:\n%s", contexts, input)
	} else {
		prompt = input
	}

	result, err := c.oa.Responses.New(ctx, responses.ResponseNewParams{
		Model:           "gpt-4.1-nano",
		MaxOutputTokens: param.NewOpt[int64](1000),
		Instructions:    param.NewOpt("You answer using only the provided context. If the context is insufficient, say so."),
		Input:           responses.ResponseNewParamsInputUnion{OfString: openai.String(prompt)},
	})

	response = result.Output[0].Content[0].Text

	return response, err
}
