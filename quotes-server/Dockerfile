FROM golang:latest

WORKDIR /app
COPY . .

RUN go build -o quotes-server .

CMD ["./quotes-server"]
