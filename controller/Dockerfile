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

# Build frontend
COPY frontend/ /app/frontend/
WORKDIR /app/frontend

RUN bun install && bun run build

# Build Go server
WORKDIR /app
COPY go.mod go.sum ./
COPY cmd/ ./cmd/
COPY internal/ ./internal/

RUN go mod download && go build -o main ./cmd/main.go

# =============================================================================
# Stage: Final runtime image

FROM debian:bookworm-slim

WORKDIR /app

# Copy the built assets
COPY --from=builder /app/main .

RUN apt-get update && \
    apt-get install -y ca-certificates curl unzip && \
    rm -rf /var/lib/apt/lists/*

# Install chrome-headless-shell
RUN curl -fsSL https://storage.googleapis.com/chrome-for-testing-public/134.0.6998.165/linux64/chrome-headless-shell-linux64.zip -o chrome.zip && \
    mkdir /chrome && \
    unzip chrome.zip -d /chrome && \
    rm chrome.zip
ENV PATH="/chrome/chrome-headless-shell-linux64:${PATH}"

ENV PIDISPLAY_HOST=0.0.0.0
ENV PIDISPLAY_PORT=3000

EXPOSE 3000

CMD ["./main"]
