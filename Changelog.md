## Create Microservices Framework

### Design Decisions / Assumptions

Starting with the Swagger API definition, which provides great base to document the API, 
lets use some off the shelf tools to work with that as much as possible.

The specification defines a number of REST endpoints that need to be delivered, however
we are not strictly limited to just REST.

Using gRPC + protobuf, we can setup a toolchain that enables changes to the Swagger API 
spec to apply some automation to the Go code. In this case, Im thinking :

Swagger / OpenAPI -> generate Protobuf -> Hand code the Go RPC endpoints.

I can then use the standard gRPC tooling to wrap the endpoints in equivalent REST tools
using https://github.com/grpc-ecosystem/grpc-gateway

(Im comfortable with that, as Ive used it in personal projects to good effect)

Benefit of this approach :

- End result is callable from both REST and RPC
- Use of protobufs applies a layer of validation
- End result is easier to change as requirements change
- End result can be more testable than hand written REST calls using a 'web framework'
