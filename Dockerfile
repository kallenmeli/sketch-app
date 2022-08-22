#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./main.go

#final stage
FROM alpine:latest
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=sketch-app Version=0.0.1
EXPOSE 8080
