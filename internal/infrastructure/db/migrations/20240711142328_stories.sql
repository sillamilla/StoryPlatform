-- +goose Up
-- +goose StatementBegin
create table stories
(
    id         text not null unique,
    user_id    text not null references users(id),
    author     text references users(username),
    title      text not null,
    text       text not null,
    rating     integer not null default 0,
    created_at timestamp not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS stories;
-- +goose StatementEnd
