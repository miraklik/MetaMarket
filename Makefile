build:
	@go build -o bin/blocker

run:
	@./bin/docker

test:
	@go test -v ./blockchain/crypto/.