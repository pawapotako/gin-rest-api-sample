####################################
## Build the application from source
####################################
FROM golang:alpine AS builder

# WORKDIR /build
# COPY . ./
# RUN go mod download
# RUN go build -o main main.go

WORKDIR /app
# ENTRYPOINT [ "tail" ]
CMD ["go","run","main.go"]

####################################
## Deploy the application binary into a lean image
####################################
# FROM alpine:latest as final
# RUN apk add --no-cache tzdata

# WORKDIR /app
# # COPY --from=builder /build/main /app/main

# CMD ["/app/main"]
