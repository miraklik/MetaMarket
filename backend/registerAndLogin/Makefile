build:
	@go build -o bin/registerAndLogin cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/registerAndLogin

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
