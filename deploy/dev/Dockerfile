FROM cosmtrek/air:v1.15.1

ENV GOPROXY=https://goproxy.io

RUN mkdir /www && mkdir /www/log

WORKDIR /news

ENTRYPOINT ["air", "-c", "/www/.air.toml"]