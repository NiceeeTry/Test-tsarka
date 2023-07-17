FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN apk add build-base && go mod download
RUN go build -o test ./cmd/api
# RUN docker run -d --name counter -p 6379:6379 --rm redis
FROM alpine:3.6
WORKDIR /app
COPY --from=builder /app .
EXPOSE 8080
CMD ["./test"]
