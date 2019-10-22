all: build

generate:
	$(MAKE) -C proto

test:
	ineffassign .
	go vet ./...
	golint ./...
	go test ./...


build: generate test
	go build

docker: generate test
	mkdir -p tmp-docker
	cp Dockerfile tmp-docker
	GOOS=linux GOARCH=amd64 go build -o tmp-docker/entrypoint
	docker build -t petstore/v1 ./tmp-docker
	@rm -rf tmp-docker
	docker images petstore/v1

