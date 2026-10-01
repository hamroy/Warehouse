# syntax=docker/dockerfile:1

# ---------- build ----------
FROM golang:1.23-alpine AS build

WORKDIR /src

# cache dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# build a static binary
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/warehouse-server ./cmd/server

# ---------- runtime ----------
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=build /out/warehouse-server /app/warehouse-server

# migrations are embedded in the binary; no SQL files needed at runtime
EXPOSE 8080

ENV PORT=8080

ENTRYPOINT ["/app/warehouse-server"]
