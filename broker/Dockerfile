FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o broker .

EXPOSE 3000

CMD [ "./broker" ]