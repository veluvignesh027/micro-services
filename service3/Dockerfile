FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o service3 .

CMD [ "./service3" ]