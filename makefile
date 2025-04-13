DB_URL=postgres://postgres:password@localhost:5432/goledger?sslmode=disable

migrate-up:
	migrate -path internal/db/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path internal/db/migrations -database "$(DB_URL)" down

psql:
	psql $(DB_URL)
