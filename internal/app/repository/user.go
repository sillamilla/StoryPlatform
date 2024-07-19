package repository

import (
	"StoryPlatforn_GIN/internal/domain/model"
	"context"
	"github.com/jackc/pgx/v5"
)

type User interface {
	Create(ctx context.Context, user model.User) error
	IsUsernameExist(ctx context.Context, userName string) (bool, error)
	GetByUsername(ctx context.Context, username string) (model.User, error)
	GetByID(ctx context.Context, id string) (model.User, error)
	GetUserIdBySession(ctx context.Context, session string) (string, error)
}

type user struct {
	pgx *pgx.Conn
}

func NewUser(pgx *pgx.Conn) User {
	return &user{pgx: pgx}
}

func (u user) Create(ctx context.Context, user model.User) error {
	query := `INSERT INTO users (id, username, password, created_at) VALUES ($1, $2, $3, $4)`
	_, err := u.pgx.Exec(ctx, query, user.ID, user.Username, user.Password, user.CreatedAt)

	return err
}

func (u user) IsUsernameExist(ctx context.Context, userName string) (bool, error) {
	query := `SELECT COUNT(*) FROM users WHERE username = $1`
	var count int
	err := u.pgx.QueryRow(ctx, query, userName).Scan(&count)

	return count > 0, err
}

func (u user) GetByUsername(ctx context.Context, username string) (model.User, error) {
	var user model.User
	query := `SELECT id, username, password, created_at FROM users WHERE username = $1`
	err := u.pgx.QueryRow(ctx, query, username).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)

	return user, err
}

func (u user) GetByID(ctx context.Context, id string) (model.User, error) {
	var user model.User
	query := `SELECT id, username, password, created_at FROM users WHERE id = $1`
	err := u.pgx.QueryRow(ctx, query, id).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)

	return user, err
}

func (u user) GetUserIdBySession(ctx context.Context, session string) (string, error) {
	var id string
	//todo dont like that it here
	query := `SELECT user_id FROM sessions WHERE session = $1`
	err := u.pgx.QueryRow(ctx, query, session).Scan(&id)

	return id, err
}
