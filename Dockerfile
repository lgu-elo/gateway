FROM alpine:latest

WORKDIR /app
COPY . /app

RUN apk add gcompat

EXPOSE 8080

ENTRYPOINT [ "/app/bin/gateway" ]
