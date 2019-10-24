#!/bin/sh
# using https://github.com/fullstorydev/grpcurl
grpcurl -plaintext -d "{\"pet_id\":${1}}" localhost:8081 petstore.PetstoreService/GetPetByID
