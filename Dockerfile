########## Builder ##########
FROM golang:1.19-alpine AS builder

# Install the latest version of Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Copy local source
COPY . /build
WORKDIR /build

RUN go build ./cmd/app.go

# # Build the binary
# RUN go build ./cmd/example/
########## Runtime ##########
FROM alpine:3

WORKDIR /

# Copy binaries from builder, the builder is just a downloaded image during the building process
COPY --from=builder /go/bin/dlv /
COPY --from=builder /build /


# Expose both the apps port and debugger
EXPOSE 8080 40000

CMD ./dlv --listen=:40000 --headless=true --api-version=2 --log exec ./app