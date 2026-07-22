package main

import (
	"fmt"
	"golang-embeddings-pinecone/cmd/config"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var err error
	var router = config.Load()

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
