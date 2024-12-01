FROM golang:1.20-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o category-service ./cmd/main.go

FROM alpine:3.14

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/category-service .

COPY .env .env

EXPOSE 50053

CMD ["./category-service"]
