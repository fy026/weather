FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

COPY conf conf
COPY service service

CMD ["./service"]