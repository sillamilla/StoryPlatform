-- +goose Up
-- +goose StatementBegin
create table users
(
    id         text not null primary key,
    username   text not null unique,
    password   text not null,
    created_at timestamp not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
