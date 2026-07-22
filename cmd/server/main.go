package main

import (
	"context"
	"fmt"
	"golang-embeddings-pinecone/cmd/config"
	"golang-embeddings-pinecone/cmd/server/providers"
	openai_provider "golang-embeddings-pinecone/cmd/server/providers/openai"
	pinecone_provider "golang-embeddings-pinecone/cmd/server/providers/pinecone"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var err error
	var router = config.Load()

	ctx := context.Background()

	pc, err := pinecone_provider.LoadPinecone(ctx)
	oa := openai_provider.LoadOpenAI()

	if err != nil {
		log.Fatalf("boot pinecone: %v", err)
	}

	providers.Container = &providers.Providers{
		Pinecone: pc,
		OpenAI:   oa,
	}

	var server = &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}

	if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
