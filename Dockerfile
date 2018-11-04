FROM python:alpine

LABEL vendor="Crypto Currency Exchange s.r.o."
LABEL maintainer="Crypto Currency Exchange s.r.o."
LABEL app="batcher-rates-coinbase"

RUN mkdir /app
WORKDIR /app

COPY ./requirements.txt /app
COPY ./*.py /app/

RUN pip install --upgrade pip \
    && pip install -r requirements.txt \
    && chmod +x /app/*.py

ENV TZ=Europe/Prague
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT "/app/app.py"
