FROM golang:1.25.1-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080

ENTRYPOINT ["./main"]
