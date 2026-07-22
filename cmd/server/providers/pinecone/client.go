package pinecone_provider

import (
	"context"
	"fmt"

	"github.com/pinecone-io/go-pinecone/v4/pinecone"
	"google.golang.org/protobuf/types/known/structpb"
)

type Config struct {
	APIKey    string
	IndexName string
	Cloud     string
	Region    string
	Namespace string
	Dimension int32
}

type Vector struct {
	ID       string
	Values   []float32
	Metadata map[string]any
}

type Match struct {
	ID       string
	Score    float32
	Values   []float32
	Metadata map[string]any
}

type Client struct {
	pc        *pinecone.Client
	index     *pinecone.IndexConnection
	indexName string
	namespace string
}

func (c *Client) Upsert(ctx context.Context, vectors []Vector) error {
	items := make([]*pinecone.Vector, 0, len(vectors))

	for _, v := range vectors {
		values := v.Values
		metadata, err := structpb.NewStruct(v.Metadata)

		if err != nil {
			return fmt.Errorf("build metadata for vector %q: %w", v.ID, err)
		}

		items = append(items, &pinecone.Vector{
			Id:       v.ID,
			Metadata: metadata,
			Values:   &values,
		})
	}

	_, err := c.index.UpsertVectors(ctx, items)

	if err != nil {
		return fmt.Errorf("upsert vectors: %w", err)
	}

	c.index.UpsertVectors(ctx, items)

	return nil
}

func (c *Client) Query(ctx context.Context, vectors []float32, topk uint32) ([]Match, error) {
	var err error
	var response []Match

	result, err := c.index.QueryByVectorValues(ctx, &pinecone.QueryByVectorValuesRequest{
		Vector:          vectors,
		TopK:            topk,
		IncludeValues:   true,
		IncludeMetadata: true,
	})

	if err != nil {
		err = fmt.Errorf("query vectors: %w", err)
		return response, err
	}

	response = make([]Match, 0, len(result.Matches))

	for _, m := range result.Matches {
		var matchValues []float32

		if m.Vector.Values != nil {
			matchValues = *m.Vector.Values
		}

		response = append(response, Match{
			ID:       m.Vector.Id,
			Score:    m.Score,
			Values:   matchValues,
			Metadata: m.Vector.Metadata.AsMap(),
		})
	}

	return response, err
}
