# build
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# runtime
FROM debian:12-slim

COPY --from=builder /cmd/server /server
ENTRYPOINT ["/server"]