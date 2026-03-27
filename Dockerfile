# build stage
FROM golang:1.26-alpine AS builder
WORKDIR /usr/src/app
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /usr/src/app/main .
EXPOSE 8080
CMD ["./main"]
