# Start from the official Golang image
FROM golang:alpine as builder
WORKDIR /chatbackend
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /chatbackend .


FROM alpine:latest
RUN apk --no-cache add ca-certificates mysql-client
WORKDIR /root/
COPY --from=builder /chatbackend/ .
COPY init.sh .
EXPOSE 8080
CMD ["./chatbackend"]
