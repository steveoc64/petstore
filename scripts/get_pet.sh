#!/bin/sh
echo USAGE: ./get_pet.sh ID
echo Getting Pet $1

curl -v -X GET "localhost:8080/pet/${1}" \
  -H  "accept: application/json" \
  -H  "Content-Type: application/json"
