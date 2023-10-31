FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o service2 .

CMD [ "./service2" ]