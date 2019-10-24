## Curl Scripts

A collection of curl scripts to hit the petstore from the command line.

- add_pet.sh ID
- add_pet_rpc.sh ID ... get the pet using direct RPC call
- delete_pet.sh
- get_pet.sh
- update_pet_form.sh

Note that the above functions are already covered by unit tests in the go code, 
but these scripts add some (possibly handy) command line options too.

eg:

`get_pet.sh NNN | jq`  to quickly see the contents of a pet in the DB
