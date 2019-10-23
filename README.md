# Petstore Microservice Implementation in Go - Oct 2019

## Project Goals

- Use the provided swagger API definition https://petstore.swagger.io
- Write it in Go, demonstrate a Microservices architecture
- One endpoint fully operational
- Time limit 4-8 hrs  (ie - just do whats needed, dont go overboard)
- Use the information provided, document decisions and assumptions
- Prove testablility
- Include Docker build files / build docker container / test vs docker container
- Show decisions / commit often


## Definition of `endpoint` vs `resource`

For the purpose of this exersize, I will be implementing the `/pet` endpoint as
defined here https://petstore.swagger.io/#/pet

There are 7 `resources` serviced by this endpoint.
- Find Pet by ID
- Update Pet - form data
- Delete Pet
- Upload an image for a pet
- Add a pet
- Update a Pet 
- Find Pets by status

## Using the makefile

There is a makefile provided to run development tasks easily.  It assumes that
you are running on a Unix-like system (Mac / Linux / BSD, etc)

- `make` runs tests, builds the code and builds the docker container. 

- `make test` to run tests
- `make build` to build the application
- `make run` to build and run locally from the command line
- `make docker` to build the docker container

- `make generate` to re-generate protobuf glue code, should the protobufs change
(requires the protobuf/grpc/grpc-gateway tools to be installed and configured)

## Curl Scripts

There are some utility curl scripts in the `./scripts` directory

- `add_pet.sh` ID NAME CATEGORY_NAME << to add a new pet
- `get_pet.sh` ID << to get a pet by ID
- `update_pet.sh` ID NAME STATUS << to update a pet
- `delete_pet.sh` API_KEY ID << to delete a pet.  The API_KEY must match the service's key. Default is ABC123

## Design Decisions - Architecture

The swagger specification defines a number of REST endpoints that need to be delivered,
however we are not strictly limited to just REST.

Define the models using standard protobuf tooling and gRPC to wrap the endpoints in
equivalent REST tools using https://github.com/grpc-ecosystem/grpc-gateway

Benefits of using gRPC for this Microservice :

- End result is callable from both REST and RPC
- Use of protobufs applies a layer of validation and control
- End result is easier to change as requirements change, as the protobuf and rpc definitions
are in 1 place, and should align with the swagger API definition.
- Adding RPC support provides a more efficient means of communicating between services
without the need to exersize REST endpoints and JSON marshalling.
- End result is easier to test compared to REST calls using a pure web framework.

## Design Decisions - Tools

Due to the use of gRPC/Protobuf, I am using the standard protobuf generators, to create the intermediate files to 
connect handler code to protobuf definitions and REST/JSON conversion.

The generated files in this case are committed as part of the repository:
- `proto/petstore.pb.go`
- `proto/petstore.pb.gw.go`

If the API changes, or there are changes needed to the protobuf file, then the developer
must install the protobuf tools, grpc generator, and grpc gateway tools, as described here :
https://github.com/grpc-ecosystem/grpc-gateway

## Critical Files to look at to implement API changes

The core files that encode the API, and therefore need to be changed
if the API changes are as follows :

- `proto/petstore.proto` Formal definitions of the models and parameters, defined in 1 place. The following `proto/petstore.pb.*` 
are generated code that need to be re-generated each time the protobuf file is altered. (see `make generate`)
- `proto/rest.yaml` This file defines the connection between the REST endpoints and the Go objects
- `handler/pets.go` The code for each of the endpoints in the "pets" range. 1 endpoint = 1 function
- `handler/pets_test.go` The test code for each of the endpoints in the "pets" range. 1 endpoint = 1 test

Other files such as `main.go`, `handler/petstore.go` defines the handler, sets the storage interface, and sets 
up all the network listeners.  This is common / boilerplate code, and need not be changed as the API
changes.

Likewise, the full end-to-end test case in `handler/petstore_test.go` provides test coverage for actually
spinning up the RPC and REST handlers, so does not need attention if the API expands.

## Notes on out of scope possibilities with more code generation

For the purpose of this exersize, I will avoid the use of any additional code generation tools
and write the implementations manually. I think that is the best way to get a result
out in the given timeframe, and demonstrate a good understanding of the basic principles
at the same time.

There is the possibility of further automating this workflow using off the shelf tools.

Using gRPC + protobuf, we can setup a toolchain that enables changes to the Swagger API 
spec to apply some automation to both the protobuf definitions and the Go code.

It is also possible to go the other way, so use protobufs and rpc definitions to 
describe the project, and then generate swagger / OpenAPI from that. 

There are existing tools such as gnostic https://github.com/googleapis/gnostic

.. which if configured correctly could enable generation of the both the protobuf
definitions, as well as scaffold out the API Go code for both client and server
sides of the service directly from the swagger files.

There is also the https://github.com/nytimes/openapi2proto project, which uses as 
more direct approach with its own custom parsers.  Looks promising to use, or possibly
as a base to build some custom code generation tooling from.

Sounds like fun, but is out of scope for this project.

## Design Decisions - Data Storage

There is no explicit requirement to implement persistent storage as part of the
microservice. However, it needs some form of storage to function (so that inserts and updates 
create some data that can be exersized by the get calls)

In a production environment where this Microservice is deployed, there may be a shared SQL
instance that is used for the storage, or there may be some other means.

This code includes the following pluggable database drivers:

- `memory` for in-memory data storage with no persistence 
- `mysql` for connecting to a mysql database
- `testdb` for an in-memory data storage with seeded data for running test cases

For the container running the petstore, the driver and DSN details can set using ENV vars.


eg - run the program with
`DATABASE=MEMORY ./petstore`  runs up the service, using an empty in-memory DB

`DATABASE=TESTDB ./petstore`  runs up the service, using an in-memory DB with half a dozen pets seeded

## Design Decisions - Docker

There are a LOT of different options when building a docker container, from including a full OS
and building the app inside the container, through to a minimalist container. 

I have chosen the minimalist approach with this project:

- The Go binary, being statically linked, is the only object needed in the container
- The binary is named `entrypoint` in the container
- To build the container, create a temp dir to place the contents in, and build from there
- Set GOOS and GOARCH when building the entrypoint as the developer may be using something other
than a linux box
- All runtime parameters are to be passed through using ENV vars in the docker container

## Misc Code Style Notes

- Interfaces.  Defined at the consumer end, and used as input params. Avoid writing code that returns an interface
wherever practical.
- Errors. Either - return them, or log them and handle in line.  Dont do both.
- Comments. Follow lint recommendations. Add a doc.go file to annotate `go doc` results.

## GRPC gotcha - form encoded data

The update with form data resource is problematic with grpc - by default, the grpc gateway does not support
form encoded data with a POST request.

I have added a form encoding wrapper to the HTTP mux to rewrite form encoded data to JSON to get around this.


## Tech Debt - XML Output

The current code base outputs JSON encoding only, I have not implemented XML output yet.  The Swagger API
spec denotes that both should be supported to fully implement the API.

It should be a simple matter of adding the encoder to the HTTPMux, but I havent had time to investigate
this yet.

## Tech Debt - Photo URLs and photo storage

Looking at the model for the Photo URLs attached to the Pet, it wasnt clear what the structure is 
meant to be. Would need to clarify that with the product owner.

Im seeing this :

```
	[
xml: OrderedMap { "wrapped": true }
string
xml: OrderedMap { "name": "photoUrl" }
xml:
   name: photoUrl]
```

For the sake of simplicity, I am implementing the PhotoURLs is a simple []string{} against the pet.

If a more complex structure is needed (map[string]somePhotoStruct ??), then its easy enough to change.

I have not implemented a persistent storage mechanism for the image data inside the petstore microservice.

In a real application, would decouple that storage entirely, so that the actual images would be stored
in a separate service (AWS S3, etc).  Depending on how that is done, that would drive how the PhotoURLs
would look.


## Tech Debt / TODO

Implementing this using gRPC/REST was pretty straight forward for the most part, but I did manage
to paint myself into a corner with the error handling.

The grpc-gateway has default fallbacks for managing errors (such as validation and parsing of input),
which is nice and consistent and saves a lot of boilerplate code.

However, the API spec does have requirements for error returns that grpc-gateway does not 
exactly align with, so I have had to consider coming up with a way to control return values
and HTTP status codes from within the handler, and have these applied by grpc-gateway.

Have added a few extra lines of code in `handler/petstore.go` to implement a custom HTTPError 
handler to enable this. Easy enough, but it does make the code more complicated.

Using go1.13 error wrapping would be an ideal solution for passing more detailed errors through, although
the grpc-gateway code does not support this yet. Nor could it without breaking compatibility for earlier
versions of Go.

Since complex error values are lost when passing through the gateway code, I have implemented a 
convention where the handler can return an error value of the form :

`NNN:Error String`

If that pattern is detected, the custom error handler then treats the leading NNN as the status code to be used,
and returns the rest as the error string. Its a valid assumption that all HTTP return codes are 3 digits.

Im not 100% happy with that approach, but it works, is easy to reason about, and makes controlling error returns
from the handlers very simple to code. Longer term, I would look at a better implementation of controlling errors
perhaps. 


