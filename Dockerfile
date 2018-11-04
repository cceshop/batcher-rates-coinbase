FROM python:alpine

LABEL vendor="Crypto Currency Exchange s.r.o."
LABEL maintainer="Crypto Currency Exchange s.r.o."

RUN mkdir /app
WORKDIR /app

RUN pip install --upgrade pip
COPY ./requirements.txt /app
RUN pip install -r requirements.txt

COPY ./*.py /app/
RUN chmod +x /app/*.py

ENV TZ=Europe/Prague
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT "/app/app.py"
