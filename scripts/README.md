## Curl Scripts

A collection of curl scripts to hit the petstore from the command line.

- `add_pet.sh` ID NAME CATEGORY_NAME << to add a new pet
- `get_pet.sh` ID << to get a pet by ID
- `update_pet.sh` ID NAME STATUS << to update a pet
- `delete_pet.sh` API_KEY ID << to delete a pet.  The API_KEY must match the service's key. Default is ABC123
- `find_by_status.sh` STATUS << to lookup pets by status code.
- `upload_file.sh` uploads the "avatar.png" sample image to test file uploads

Note that the above functions are already covered by unit tests in the go code, 
but these scripts add some (possibly handy) command line options too.

eg:

`get_pet.sh NNN | jq`  to quickly see the contents of a pet in the DB
