-- +goose Up
-- +goose StatementBegin
CREATE TABLE story_ratings
(
  user_id text not null references users(id) on delete cascade,
  story_id text not null references stories(id) on delete cascade,

  UNIQUE (user_id, story_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists story_ratings;
-- +goose StatementEnd
