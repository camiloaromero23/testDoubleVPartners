FROM ubuntu:20.04

WORKDIR /app

COPY main /app

ENV PORT 8081

ENV DB_HOST localhost

ENV DB_PORT 3306

EXPOSE $PORT

CMD ["./main"]