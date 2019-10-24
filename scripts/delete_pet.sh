#!/bin/sh

echo These curl scripts will hit a running test server with a valid pet ID = 5
echo and try some combinations of deleting a pet.
echo
echo You can run the server locally, and set DATABASE=TESTDB ./petstore
echo to have a have a valid set of test data
echo

echo Delete a valid pet without an api key == fail
echo hit any key
read anykey
curl -v -X DELETE -i http://localhost:8080/pet/5
echo

echo delete valid pet with invalid key == fail
echo hit any key
read anykey
curl -v -X DELETE -H 'api_key: ABCxxx' -i http://localhost:8080/pet/5
echo

echo delete invalid pet with valid key == fail
echo hit any key
read anykey
curl -v -X DELETE -H 'api_key: ABC123' -i http://localhost:8080/pet/100
echo

echo delete valid pet with valid key == pass
echo hit any key
read anykey
curl -v -X DELETE -H 'api_key: ABC123' -i http://localhost:8080/pet/5
echo

