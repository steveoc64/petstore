#!/bin/sh
echo USAGE: ./add_pet.sh ID NAME CATEGORY
echo Adding Pet $1 with name $2 of category $3

curl -v -X POST "localhost:8080/pet" \
  -H  "accept: application/json" \
  -H  "Content-Type: application/json" \
  -d "{\"id\": ${1}, \
    \"category\": {\"id\":1,\"name\":\"${3}\"}, \
    \"name\": \"${2}\", \
    \"photoUrls\":[\"myphotos.com/1234.jpg\"], \
    \"tags\":[{\"id\":1,\"name\":\"housetrained\"},{\"id\":2,\"name\":\"Likes Kids\"}], \
    \"status\":\"available\"}"

echo
