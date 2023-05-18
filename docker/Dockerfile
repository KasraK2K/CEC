FROM alpine:latest

RUN mkdir /app
COPY ./backend /app/

RUN mkdir -p /pkg/config
COPY ./.env /pkg/config/

CMD [ "/app/backend" ]