FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o service3 .

EXPOSE 3003

CMD [ "./service3" ]