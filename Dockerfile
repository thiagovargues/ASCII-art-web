# syntax=docker/dockerfile:1

FROM golang:1.23-alpine AS builder
WORKDIR /app

# Cache modules and build artifacts
RUN --mount=type=cache,target=/root/.cache/go-build true

COPY go.mod ./
COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/ascii-art-web .

FROM alpine:3.20
WORKDIR /app

# Non-root runtime user
RUN addgroup -S asciiart && adduser -S asciiart -G asciiart

COPY --from=builder /app/ascii-art-web /app/
COPY templates/ /app/templates/
COPY standard.txt shadow.txt thinkertoy.txt style.css /app/

ENV PORT=8080
EXPOSE 8080

USER asciiart

LABEL org.opencontainers.image.title="ascii-art-web" \
      org.opencontainers.image.description="ASCII art web server in Go" \
      org.opencontainers.image.authors="Alioune Sall, Emilia Chedot, Thiago Vargues" \
      org.opencontainers.image.licenses="MIT" \
      org.opencontainers.image.version="latest"

ENTRYPOINT ["/app/ascii-art-web"]
