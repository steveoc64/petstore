#!/bin/sh

curl -X POST "http://localhost:8080/pet/1/uploadImage" -H  "accept: application/json" -H  "Content-Type: multipart/form-data" -F "file=@avatar.png;type=image/png"
echo
