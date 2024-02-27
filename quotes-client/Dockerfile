FROM golang:latest

WORKDIR /app
COPY . .

RUN go build -o quotes-client .

CMD ["./quotes-client"]
