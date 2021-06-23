FROM golang:1.16

COPY /main .
COPY /.env .
COPY /config/config.toml .

CMD ./main