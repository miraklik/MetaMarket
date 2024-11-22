run:
	go run cmd/main.go

migrations:
	migrate create -ext sql -dir db/migrations/ -seq init_marketdb 

migrate-up:
	go run cmd/main.go migrate up

migrate-down:
	go run cmd/main.go migrate down

.PHONY: run migrate-up migrate-down