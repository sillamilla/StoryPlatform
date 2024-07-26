

migrate-up:
	goose -dir ./internal/infrastructure/db/migrations postgres "postgres://postgres123:postgres123@localhost:5432/postgres123?sslmode=disable" up

migrate-down:
	goose -dir ./internal/infrastructure/db/migrations postgres "postgres://postgres123:postgres123@localhost:5432/postgres123?sslmode=disable" down

migrate-down-to:
	goose -dir ./internal/infrastructure/db/migrations create new_user_table sql

test.unit:
	go test --short ./...

test.integration:
	docker run --name tests-postgres -e POSTGRES_USER=test1234 -e POSTGRES_PASSWORD=test1234 -e POSTGRES_DB=test1234 -p 5432:5432 -d postgres:13
	goose -dir ./internal/infrastructure/db/migrations postgres "postgres://test1234:test1234@localhost:5432/test1234?sslmode=disable" up
	$env:GIN_MODE="release"; go test -v ./tests/.

	docker stop tests-postgres
	docker rm tests-postgres
