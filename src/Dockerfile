FROM golang:1.11.11

ADD . /app
WORKDIR /app

RUN mkdir /verenav
RUN go build -v

CMD ["./yurararan"]
