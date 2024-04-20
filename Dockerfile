FROM golang:1.21.0-alpine AS builder

RUN apk add --no-cache \
    gcc \
    musl-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN GOOS=linux CGO_ENABLED=1 GOARCH=amd64 \
    go build -ldflags="-w -s" -o main -tags timetzdata

FROM alpine

WORKDIR /app
COPY --from=builder /app/main ./main

# Install sqlite driver (optional, depending on your image)
RUN apk add --no-cache sqlite

EXPOSE 8080
CMD ["./main"]
