# golang-embeddings-pinecone

Small Go learning project for:

- generating embeddings with OpenAI
- storing vectors in Pinecone
- retrieving related text as context for question answering

The project uses `gin`, `docker compose`, OpenAI embeddings/responses, and Pinecone as the vector store.

## Requirements

- Go `1.25`
- Docker and Docker Compose
- OpenAI API key
- Pinecone API key

## Environment

Create a local `.env` file based on `.env.example`.

Example:

```env
PORT=3000

GOFLAGS=-tags=nomsgpack

PINECONE_API_KEY=your-pinecone-api-key
PINECONE_INDEX=embeddings
PINECONE_CLOUD=aws
PINECONE_REGION=us-east-1
PINECONE_NAMESPACE=default
PINECONE_DIMENSION=1536

OPENAI_API_KEY=your-openai-api-key
GEMINI_API_KEY=your-gemini-ai-key
```

## Run With Docker

Start the dev container:

```sh
docker compose up -d
```

Run the Go server inside the container:

```sh
docker compose exec server go run ./cmd/server
```

The API will be available on:

```txt
http://localhost:3000
```

## API

### Health

```http
GET /api/v1/health/ping
```

### Embeddings

```http
POST /api/v1/embeddings/create
Content-Type: application/json
```

Body:

```json
{
  "text": "your text to embed"
}
```

### Retrieval

```http
POST /api/v1/retrieval/ask
Content-Type: application/json
```

Body:

```json
{
  "input": "your question"
}
```

## Project Structure

```txt
cmd/
  config/
  server/
    api/
      health/
      embeddings/
      retrieval/
    providers/
      openai/
      pinecone/
```

## Notes

- Pinecone index dimension is expected to match the embedding model.
- With `text-embedding-3-small`, the dimension is `1536`.
- `.env` is ignored by git.
- Rotate credentials immediately if they were ever exposed.
