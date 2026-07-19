DB_URL=postgres://ecom_user:ecom_pass@localhost:5432/ecom?sslmode=disable

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

run:
	go run ./cmd/api
