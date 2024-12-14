run:
	go run cmd/main.go

migrations:
	migrate create -ext sql -dir db/migrations/ -seq init_marketdb 

migrate-up:
	go run cmd/main.go migrate up

migrate-down:
	go run cmd/main.go migrate down

docker:
	docker-compose up --build

.PHONY: run migrations migrate-up migrate-down, docker