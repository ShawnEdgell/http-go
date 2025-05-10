FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /http-go-app ./main.go

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /http-go-app .
EXPOSE 8080
CMD ["./http-go-app"]