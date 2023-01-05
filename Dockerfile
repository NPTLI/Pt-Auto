FROM golang:1.16-alpine as builder

WORKDIR /pt-auto/server
COPY . .

RUN cd server \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="pt-auto@apt@pt.com"

WORKDIR /pt

COPY --from=0 /pt-auto/server/server ./

EXPOSE 666
ENTRYPOINT ./server


