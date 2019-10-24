#!/bin/sh

curl "http://localhost:8080/findByStatus?status=available&status=sold" -H  "accept: application/json"
echo
