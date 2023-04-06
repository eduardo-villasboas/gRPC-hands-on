FROM golang

WORKDIR /home/

RUN go install github.com/ktr0731/evans@latest

RUN apt update && apt install sqlite3 -y

