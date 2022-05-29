FROM golang:1.18.1-alpine3.15 AS builder

RUN apk add --update --no-cache ca-certificates
RUN apk add --update make
RUN apk add git

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image
# and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN make publish

RUN touch .env

FROM alpine:3.13.6

# Copy binary and config files from /build
# to root folder of scratch container.
COPY --from=builder ["/build/apiserver", "/"]

# empty env file
COPY --from=builder ["/build/.env", "/"]

# Export necessary port.
EXPOSE 9090

# Command to run when starting the container.
ENTRYPOINT ["/apiserver"]
