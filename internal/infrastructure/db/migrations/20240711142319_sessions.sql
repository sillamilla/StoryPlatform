-- +goose Up
-- +goose StatementBegin
create table sessions
(
    user_id    text not null unique references users(id),
    session    text not null,
    created_at timestamp not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sessions;
-- +goose StatementEnd
