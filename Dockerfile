#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./main.go

#final stage
FROM alpine:latest
WORKDIR /app
COPY pages ./pages
COPY --from=builder /go/bin/app /app/run
ENTRYPOINT /app/run
LABEL Name=sketch-app Version=0.0.1
EXPOSE 8080
