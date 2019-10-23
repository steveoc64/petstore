#!/bin/sh

echo "Pet 1 looks like this"
curl -X GET "localhost:8080/pet/1" -H  "accept: application/json" -H  "Content-Type: application/json"

echo
echo "Updating pet 1 using form encoded data and a POST at the same time"
curl -X POST -H 'Content-Type: application/x-www-form-urlencoded' -i http://localhost:8080/pet/1 --data 'name=FormName&status=pending'
echo

echo "Pet 1 now looks like this"
curl -X GET "localhost:8080/pet/1" -H  "accept: application/json" -H  "Content-Type: application/json"
echo


