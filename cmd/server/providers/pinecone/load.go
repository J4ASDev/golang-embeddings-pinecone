package pinecone_provider

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/pinecone-io/go-pinecone/v4/pinecone"
)

func LoadPinecone(ctx context.Context) (*Client, error) {
	var pc *pinecone.Client
	var err error

	dimension, err := strconv.ParseInt(os.Getenv("PINECONE_DIMENSION"), 10, 32)

	if err != nil {
		return nil, fmt.Errorf("parse PINECONE_DIMENSION: %w", err)
	}

	var config = Config{
		APIKey:    os.Getenv("PINECONE_API_KEY"),
		IndexName: os.Getenv("PINECONE_INDEX"),
		Cloud:     os.Getenv("PINECONE_CLOUD"),
		Region:    os.Getenv("PINECONE_REGION"),
		Namespace: os.Getenv("PINECONE_NAMESPACE"),
		Dimension: int32(dimension),
	}

	pc, err = pinecone.NewClient(pinecone.NewClientParams{
		ApiKey: config.APIKey,
	})

	if err != nil {
		return nil, err
	}

	if err := ensureIndex(ctx, pc, config); err != nil {
		return nil, err
	}

	desc, err := pc.DescribeIndex(ctx, config.IndexName)

	if err != nil {
		err = fmt.Errorf("describe index %q: %w", config.IndexName, err)
		return nil, err
	}

	idxConnection, err := pc.Index(pinecone.NewIndexConnParams{
		Host: desc.Host,
	})

	if err != nil {
		return nil, fmt.Errorf("connect index %q: %w", config.IndexName, err)
	}

	return &Client{
		pc:        pc,
		index:     idxConnection,
		indexName: config.IndexName,
		namespace: config.Namespace,
	}, err
}

func ensureIndex(ctx context.Context, pc *pinecone.Client, cfg Config) error {
	indexes, err := pc.ListIndexes(ctx)

	if err != nil {
		return fmt.Errorf("list indexes: %w", err)
	}

	for _, idx := range indexes {
		if idx.Name == cfg.IndexName {
			return waitUntilReady(ctx, pc, cfg.IndexName)
		}
	}

	metric := pinecone.Cosine
	dimension := cfg.Dimension

	_, err = pc.CreateServerlessIndex(ctx, &pinecone.CreateServerlessIndexRequest{
		Name:      cfg.IndexName,
		Dimension: &dimension,
		Metric:    &metric,
		Cloud:     pinecone.Cloud(cfg.Cloud),
		Region:    cfg.Region,
	})

	if err != nil {
		return fmt.Errorf("create index %q: %w", cfg.IndexName, err)
	}

	return waitUntilReady(ctx, pc, cfg.IndexName)
}

func waitUntilReady(ctx context.Context, pc *pinecone.Client, indexName string) error {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	timeout := time.NewTimer(2 * time.Minute)
	defer timeout.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-timeout.C:
			return fmt.Errorf("timeout waiting for index %q to become ready", indexName)
		case <-ticker.C:
			desc, err := pc.DescribeIndex(ctx, indexName)
			if err != nil {
				return fmt.Errorf("describe index %q while waiting: %w", indexName, err)
			}

			if desc.Status.Ready {
				return nil
			}
		}
	}
}
