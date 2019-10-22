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

## Using the makefile

There is a makefile provided to run development tasks easily.  It assumes that
you are running on a Unix-like system (Mac / Linux / BSD, etc)

- `make` runs tests, code generators, builds the code and builds the docker container

- `make generate` to execute any code generators needed
- `make test` to run tests
- `make build` to build the application
- `make run` to build and run locally from the command line
- `make docker` to build the docker container

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
- End result can be more testable than hand written REST calls using a pure web framework

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

## Design Decisions - Tools

Due to the use of gRPC/Protobuf, I am using the standard protobuf generators, to create the intermediate files to 
connect handler code to protobuf definitions and REST/JSON conversion.

The generated files in this case are committed as part of the repository:
- `proto/petstore.pb.go`
- `proto/petstore.pb.gw.go`

If the API changes, or there are changes needed to the protobuf file, then the developer
must install the protobuf tools, grpc generator, and grpc gateway tools, as described here :
https://github.com/grpc-ecosystem/grpc-gateway

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

So, I will define a storage interface as part of the petstore, and implement an 'in memory' 
storage that can be used.  This would assist in making the Microservice testable without the
need to mount dependent services (such as a Database)

In addition to the 'in memory' storage implementation, I will also implement an SQL storage
class that connects to another container to do the IO.

This only needs to be done to cover 1 API endpoint for this exersize, so shouldnt be too much
extra work.

For the container running the petstore, the driver and DSN details can set using ENV vars.

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
wherevere practical.
- Errors. Either - return them, or log them and handle in line.  Dont do both.
- Comments. Follow lint recommendations. Add a doc.go file to annotate `go doc` results.
