FROM golang:1.9-alpine
MAINTAINER i@boris.tech

COPY . /go/src/github.com/borisding1994/hathcoin
WORKDIR /go/src/github.com/borisding1994/hathcoin
RUN apk --no-cache add make openssl git

# build binary
RUN make dep-install && \
    make build && \
    mkdir -p /app

# remove sources
RUN cp -r /go/src/github.com/borisding1994/hathcoin/dist /app && \
    rm -rf /go/src/*

WORKDIR /app
#VOLUME /app/logs
EXPOSE 8081 8081
ENTRYPOINT ["/app/hathcoin", "start"]