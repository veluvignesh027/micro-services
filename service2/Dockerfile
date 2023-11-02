FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o service2 .

EXPOSE 3002

CMD [ "./service2" ]