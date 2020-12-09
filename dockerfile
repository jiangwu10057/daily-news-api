FROM golang:1.15.6-alpine3.12 as build

ENV GOPROXY=https://goproxy.io

ADD ./src /news

WORKDIR /news

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o news_server

FROM alpine:3.7 as alpine-n

RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf && \
    mkdir -p /www/log

WORKDIR /www

COPY --from=build /news/news_server /usr/bin/news_server

ADD ./conf /www/conf

RUN chmod +x /usr/bin/news_server

RUN export GIN_MODE=release

ENTRYPOINT ["news_server"]
