FROM ubuntu:18.04

WORKDIR /app

ADD wine /app
RUN chmod 544 /app/wine

EXPOSE 8500
ENTRYPOINT ["/app/wine"]