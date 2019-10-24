#!/bin/sh
curl "localhost:8080/pet/${1}" \
  -H  "accept: application/json" \
  -H  "Content-Type: application/json"

echo
