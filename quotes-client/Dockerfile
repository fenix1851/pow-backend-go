FROM golang:latest

# Install netcat-openbsd
RUN apt-get update && apt-get install -y netcat-openbsd && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY . .

RUN go build -o quotes-client .

CMD ["./quotes-client -baseUrl http://quotes-server:8080"]

