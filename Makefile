POSTGRES_USER=postgres123
POSTGRES_PASSWORD=postgres123
POSTGRES_DB=postgres123

migrate-up:
	goose -dir ./internal/infrastructure/db/migrations postgres "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable" up

migrate-down:
	goose -dir ./internal/infrastructure/db/migrations postgres "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable" down

test.unit:
	go test --short ./...

test.integration:
	goose -dir ./internal/infrastructure/db/migrations postgres "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable" up
	go test -v ./tests/.
