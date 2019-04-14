FROM golang:1.12.4-alpine3.9

LABEL cce.image.vendor="Crypto Currency Exchange s.r.o."
LABEL cce.image.maintainer="Crypto Currency Exchange s.r.o."
LABEL cce.git.repository="https://github.com/cceshop/batcher-rates-coinbase.git"
LABEL cce.quay.repository="https://quay.io/repository/cceshop/batcher-rates-coinbase"

USER root
RUN mkdir -p /app
WORKDIR /app
COPY ./entrypoint.sh /app
COPY ./get_rates.go /app
RUN chmod +x /app/entrypoint.sh \ 
    && apk add git \
    && go get github.com/go-redis/redis

ENTRYPOINT ["/app/entrypoint.sh"]
