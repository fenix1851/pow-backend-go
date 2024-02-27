#!/bin/sh

while ! nc -z quotes-server 8080; do
echo 'Waiting for the quotes-server to come online...'
sleep 1
done
echo 'Quotes-server is up and running!'
go run main.go --baseUrl http://quotes-server:8080