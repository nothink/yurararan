FROM golang:1.12.9

ADD . /app
WORKDIR /app

RUN mkdir /verenav

RUN go build -v

CMD ["./yurararan"]
