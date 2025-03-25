# Stage: Build
FROM debian:bookworm-slim AS builder

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
    apt-get install --no-install-recommends -y ca-certificates curl unzip && \
    # Chrome Deps
    apt-get install -y \
        fonts-liberation \
        libappindicator3-1 \
        libasound2 \
        libatk-bridge2.0-0 \
        libatk1.0-0 \
        libc6 \
        libcairo2 \
        libcups2 \
        libdbus-1-3 \
        libexpat1 \
        libfontconfig1 \
        libgbm1 \
        libgcc1 \
        libglib2.0-0 \
        libgtk-3-0 \
        libnspr4 \
        libnss3 \
        libpango-1.0-0 \
        libpangocairo-1.0-0 \
        libstdc++6 \
        libx11-6 \
        libx11-xcb1 \
        libxcb1 \
        libxcomposite1 \
        libxcursor1 \
        libxdamage1 \
        libxext6 \
        libxfixes3 \
        libxi6 \
        libxrandr2 \
        libxrender1 \
        libxss1 \
        libxtst6 \
        xdg-utils && \
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
