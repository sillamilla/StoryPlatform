migrate-up:
	goose -dir ./internal/infrastructure/db/migrations postgres "postgres://test1234:test1234@localhost:5432/test1234?sslmode=disable" up

migrate-down:
	goose -dir ./internal/infrastructure/db/migrations postgres "postgres://test1234:test1234@localhost:5432/test1234?sslmode=disable" down

migrate-down-to:
	goose -dir ./internal/infrastructure/db/migrations create new_user_table sql

test.unit:
	go test --short ./...

test.integration:
	goose -dir ./internal/infrastructure/db/migrations postgres "postgres://test1234:test1234@localhost:5432/test1234?sslmode=disable" up
	go test -v ./tests/.
