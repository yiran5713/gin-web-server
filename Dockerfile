# Builder stage
FROM golang:1.21-alpine as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add curl
WORKDIR /root/
COPY --from=builder /app/myapp .
CMD ["./myapp"]
