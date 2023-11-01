# syntax=docker/dockerfile:1

####################################
## Build the application from source
####################################
FROM golang:alpine AS builder

WORKDIR /build
COPY . ./
RUN go mod download
RUN go build -o main cmd/api/main.go

####################################
## Deploy the application binary into a lean image
####################################
# FROM scratch as final
# FROM busybox as final
# FROM gcr.io/distroless/base-debian10
FROM alpine:latest as final
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /build/main /app/main
COPY --from=builder /build/configs/config.yaml /app/configs/config.yaml

ENTRYPOINT ["/app/main"]
