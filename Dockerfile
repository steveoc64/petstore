# Dockerfile for the Petstore
FROM scratch
MAINTAINER steveoc64@gmail.com

# Statically linked Go code - just need the 1 executable and thats all
COPY ./entrypoint ./entrypoint
ENTRYPOINT ["/entrypoint"]

# Set the port to expose the service on
ENV REST_PORT 8080
EXPOSE 8080/tcp

ENV RPC_PORT 8081
EXPOSE 8081/tcp

# Some endpoints (like delete pet) need an API_KEY value - define it here
ENV API_KEY ABC123

# Which storage mechanism to use ?
# Use "memory" for in-memory storage
# or set an appropriate MySQL / pgSQL / etc
ENV DATABASE memory
ENV DSN none

#ENV DATABASE mysql
#ENV DSN user:password@/dbname
