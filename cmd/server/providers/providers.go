package providers

import (
	openai_provider "golang-embeddings-pinecone/cmd/server/providers/openai"
	pinecone_provider "golang-embeddings-pinecone/cmd/server/providers/pinecone"
)

type Providers struct {
	Pinecone *pinecone_provider.Client
	OpenAI   *openai_provider.Client
}

var Container *Providers
