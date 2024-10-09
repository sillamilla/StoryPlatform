migrate-up:
	goose -dir ./internal/infrastructure/db/migrations postgres "postgres://postgres123:postgres123@localhost:5432/postgres123?sslmode=disable" up

migrate-down:
	goose -dir ./internal/infrastructure/db/migrations postgres "postgres://postgres123:postgres123@localhost:5432/postgres123?sslmode=disable" down

test.unit:
	go test --short ./...

test.integration:
	goose -dir ./internal/infrastructure/db/migrations postgres "postgres://postgres123:postgres123@localhost:5432/postgres123?sslmode=disable" up
	go test -v ./tests/.

