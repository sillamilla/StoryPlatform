-- +goose Up
-- +goose StatementBegin
create table stories
(
    id         text not null unique,
    user_id    text not null  unique  references users(id),
    author     text references users(username),
    title      text not null,
    text       text not null,
    rating     integer default 0,
    created_at timestamp not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS stories;
-- +goose StatementEnd
