# Build image: docker build -t relayers/fabric .
FROM golang:1.15-alpine3.12 as builder

# Set up dependencies
ENV PACKAGES make git libc-dev bash
WORKDIR $GOPATH/src
COPY . .
# Install minimum necessary dependencies, build binary
RUN apk add --no-cache $PACKAGES
make install

FROM alpine:3.12
COPY --from=builder /go/bin/relayer /usr/local/bin/relayer

CMD ["relayer"]
