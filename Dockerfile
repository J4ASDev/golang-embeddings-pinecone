FROM golang:1.25-alpine AS base

WORKDIR /root/code

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY ["./cmd", "./cmd"]

FROM base AS builder

WORKDIR /root/code

RUN go build -o app ./cmd/server

FROM golang:1.25-alpine AS runner

WORKDIR /root/code

EXPOSE 3000
COPY --from=builder ["/root/code/app", "."]

ARG VERSION
ENV VERSION=${VERSION}

ENV PORT=3000
ENV GIN_MODE=release

ENTRYPOINT ["./app"]
