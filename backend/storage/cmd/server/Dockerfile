# build
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o health-storage ./cmd/server

# runtime
FROM debian:12-slim

# /usr/local/bin/server is good for debian-slim but not for distroless
COPY --from=builder /app/health-storage /usr/local/bin/health-storage
ENTRYPOINT ["health-storage"]