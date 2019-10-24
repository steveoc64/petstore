#!/bin/sh

curl "http://localhost:8080/pet/findByStatus?status=$1" -H  "accept: application/json"
echo
