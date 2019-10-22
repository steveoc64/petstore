all: build

generate:
	$(MAKE) -C proto

test:
	ineffassign .
	go vet ./...
	golint ./...
	go test ./...


build: test
	go build

run: build
	RPC_PORT=8081 REST_PORT=8080 API_KEY=ABC123 ./petstore

docker: test
	mkdir -p tmp-docker
	cp Dockerfile tmp-docker
	GOOS=linux GOARCH=amd64 go build -o tmp-docker/entrypoint
	docker build -t petstore/v1 ./tmp-docker
	@rm -rf tmp-docker
	docker images petstore/v1

