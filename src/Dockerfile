FROM golang:1.13.8

ADD . /app
WORKDIR /app

RUN mkdir /verenav

RUN go build -v

CMD ["./yurararan"]
