FROM alpine:latest

RUN mkdir /app
RUN mkdir -p /app/pkg/config

COPY ./backend /app/
COPY ./.env /app/pkg/config/

CMD [ "/app/backend" ]