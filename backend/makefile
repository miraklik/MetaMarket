run:
	go run cmd/main.go

migrations:
	migrate create -ext sql -dir db/migrations/ -seq init_marketdb 

migrate-up-worker:
	go run cmd/worker/main.go migrate up

migrate-down-worker:
	go run cmd/worker/main.go migrate down

migrate-up-api:
	go run cmd/api/main.go migrate up

migrate-down-api:
	go run cmd/api/main.go migrate down

docker:
	docker-compose up --build

.PHONY: run migrations migrate-up-worker migrate-down-worker docker