# Stage: Build
FROM debian:bookworm AS builder

RUN apt-get update && apt-get install -y curl git tar unzip && rm -rf /var/lib/apt/lists/*

# Install Go from the official tarball
RUN curl -fsSL https://go.dev/dl/go1.24.1.linux-amd64.tar.gz -o go.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go.tar.gz && \
    rm go.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

# Install Bun
RUN curl -fsSL https://bun.sh/install | bash -s "bun-v1.2.6"
ENV PATH="/root/.bun/bin:${PATH}"

# Set working directory and copy project files
WORKDIR /app
COPY go.mod go.sum ./
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY frontend/ ./frontend/

# Build frontend 
WORKDIR /app/frontend
RUN bun install && bun run build

# Build the Go server
WORKDIR /app
RUN ( go mod download || true )&& go build -o main ./cmd/main.go

# Stage: Final runtime image
FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the Go binary and built frontend assets from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/internal/server/dist ./dist

ENV PIDISPLAY_HOST=0.0.0.0
ENV PIDISPLAY_PORT=3000

EXPOSE 3000

CMD ["./main"]
