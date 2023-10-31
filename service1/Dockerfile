FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o service1 .

CMD [ "./service1" ]