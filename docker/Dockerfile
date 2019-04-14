FROM golang:1.12.4-alpine3.9

USER root
RUN mkdir -p /app
WORKDIR /app
COPY ./entrypoint.sh /app
COPY ./get_rates.go /app
RUN chmod +x /app/entrypoint.sh \ 
    && apk add git \
    && go get github.com/go-redis/redis

ENTRYPOINT ["/app/entrypoint.sh"]
