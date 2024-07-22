package repository

import (
	"StoryPlatforn_GIN/internal/domain/model"
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5"
)

type Story interface {
	Create(ctx context.Context, input model.Story) error
	Get(ctx context.Context, id string) (model.Story, error)

	IsRated(ctx context.Context, userID string, id string) (string, error)
	Rate(ctx context.Context, id string, rate int) error
	MarkUserRated(ctx context.Context, userID string, id string) error

	Update(ctx context.Context, userID string, id string, input model.StoryInput) error
	Delete(ctx context.Context, userID, id string) error
}

type story struct {
	pgx *pgx.Conn
}

func NewStory(pgx *pgx.Conn) Story {
	return &story{
		pgx: pgx,
	}
}

func (s story) Create(ctx context.Context, input model.Story) error {
	query := "INSERT INTO stories(id, user_id, author, title, text, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := s.pgx.Exec(ctx, query, input.ID, input.UserID, input.Author, input.Title, input.Text, input.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s story) Get(ctx context.Context, id string) (model.Story, error) {
	var data model.Story

	query := "SELECT id, user_id, author, title, text, rating, created_at FROM stories WHERE id = $1"
	err := s.pgx.QueryRow(ctx, query, id).Scan(&data.ID, &data.UserID, &data.Author, &data.Title, &data.Text, &data.Rating, &data.CreatedAt)
	if err != nil {
		return model.Story{}, err
	}

	return data, nil
}

func (s story) IsRated(ctx context.Context, userID string, id string) (string, error) {
	var data string

	query := "SELECT story_id FROM story_ratings WHERE user_id = $1 AND story_id = $2"
	err := s.pgx.QueryRow(ctx, query, userID, id).Scan(&data)
	if err != nil {
		return "", err
	}

	return data, nil
}

func (s story) Rate(ctx context.Context, id string, rate int) error {
	query := "UPDATE stories SET rating = rating + $1 WHERE id = $2"
	_, err := s.pgx.Exec(ctx, query, rate, id)
	if err != nil {
		return err
	}

	return nil
}

func (s story) MarkUserRated(ctx context.Context, userID string, id string) error {
	query := `INSERT INTO story_ratings(user_id, story_id) VALUES ($1, $2)`
	_, err := s.pgx.Exec(ctx, query, userID, id)
	if err != nil {
		return err
	}

	return nil
}

func (s story) Update(ctx context.Context, userID string, id string, input model.StoryInput) error {
	query := "UPDATE stories SET title = $1, text = $2 WHERE id = $3 AND user_id = $4"
	result, err := s.pgx.Exec(ctx, query, input.Title, input.Text, id, userID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (s story) Delete(ctx context.Context, userID string, id string) error {
	query := "DELETE FROM stories WHERE id = $1 AND user_id = $2"
	result, err := s.pgx.Exec(ctx, query, id, userID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return sql.ErrNoRows
	}

	return nil
}
