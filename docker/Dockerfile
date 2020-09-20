FROM golang:1.15.2-alpine3.12

LABEL cce.image.vendor="Crypto Currency Exchange s.r.o."
LABEL cce.image.maintainer="Crypto Currency Exchange s.r.o."
LABEL cce.git.repository="https://github.com/cceshop/batcher-rates-coinbase.git"
LABEL cce.quay.repository="https://quay.io/repository/cceshop/batcher-rates-coinbase"

ONBUILD USER root
RUN adduser --disabled-password --gecos '' -s /bin/bash cce \
    && mkdir -p /app
COPY ./entrypoint.sh /app
COPY ./get_rates.go /app
RUN chmod +x /app/entrypoint.sh \
    && apk add git tini \
    && go get github.com/go-redis/redis \
    && chown -R cce:cce /app

USER cce
WORKDIR /app

ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/app/entrypoint.sh"]
