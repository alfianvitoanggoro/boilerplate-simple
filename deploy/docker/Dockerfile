# Stage 1: Development image
FROM golang:1.24-alpine AS development

RUN apk add --no-cache git curl

# Install Air ke global path
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin

WORKDIR /app

# Copy semua project ke /app
COPY ../../ .

# Tambahkan .air.toml (pastikan ada)
COPY ../../.air.toml .air.toml

# Pastikan folder bin ada agar air bisa build
RUN mkdir -p bin

CMD ["air"]



# Stage 2: Builder image
FROM golang:1.24-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOARCH=amd64 \
    GOOS=linux

WORKDIR /build

# Copy go.mod and go.sum first
COPY go.mod go.sum ./

# (Optional) Copy .env if needed during build
COPY .env ./

# Download dependencies
RUN go mod download

# Copy all project source code
COPY . .

# Create output directory
RUN mkdir -p .dist

# Build binary ke folder .dist/
RUN go build --ldflags "-extldflags -static" -o .dist/boilerplate-simple ./cmd/server/main.go


# Stage 3: Runtime image
FROM alpine:latest

RUN apk add --no-cache tini

WORKDIR /app
COPY --from=builder /build/.env .env
COPY --from=builder /build/.dist/ /app/.dist/

EXPOSE 50051

CMD ["./.dist/boilerplate-simple"]