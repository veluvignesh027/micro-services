FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o service1 .

EXPOSE 3001

CMD [ "./service1" ]