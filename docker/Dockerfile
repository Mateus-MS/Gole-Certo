# Image to compile go
FROM golang:1.24 AS builder

WORKDIR /app

# Copy all files into container
COPY ../backend .

# Compile go
RUN go mod tidy
RUN go build -o server .

# Image that gonna run
FROM debian:bookworm-slim
WORKDIR /app

# Move the compiled go to the run-time image
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"] 