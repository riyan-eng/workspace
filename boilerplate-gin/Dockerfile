# build
FROM riyaneng/golang:1.22.3-alpine AS build-stage
WORKDIR /app
COPY . .
RUN go mod download
RUN swag init -o ./docs
RUN CGO_ENABLED=0 GOOS=linux go build -o /binary

# release
FROM alpine:3.20.1 AS release-stage
RUN wget -O go.tgz https://go.dev/dl/go1.22.3.linux-amd64.tar.gz 
RUN tar -C /usr/local -xzf go.tgz 

WORKDIR /app
COPY --from=build-stage /binary /app/binary
CMD ["/app/binary"]