#!/bin/bash

rm cache
go build -o cache ./src
# docker-compose up -d
./cache

# docker-compose down