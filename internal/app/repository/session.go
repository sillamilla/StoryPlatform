package repository

import (
	"StoryPlatforn_GIN/internal/domain/model"
	"context"
	"github.com/jackc/pgx/v5"
	"time"
)

type Session interface {
	Upsert(ctx context.Context, userID, session string, time time.Time) error
	GetUserID(ctx context.Context, session string) (string, error)
	Delete(ctx context.Context, session string) error
	GetSessionInfo(ctx context.Context, session string) (model.SessionInfo, error)
}

type sessions struct {
	pgx *pgx.Conn
}

func NewSession(pgx *pgx.Conn) Session {
	return &sessions{pgx: pgx}
}

func (s *sessions) Upsert(ctx context.Context, userID, session string, time time.Time) error {
	query := `
        INSERT INTO sessions (user_id, session, created_at) 
        VALUES ($1, $2, $3) 
        ON CONFLICT (user_id) 
        DO UPDATE SET session = $2, created_at = $3
    `
	_, err := s.pgx.Exec(ctx, query, userID, session, time)
	return err
}

func (s *sessions) GetUserID(ctx context.Context, session string) (string, error) {
	var userID string

	query := "SELECT user_id FROM sessions WHERE session = $1"
	err := s.pgx.QueryRow(ctx, query, session).Scan(&userID)
	return userID, err
}

func (s *sessions) GetSessionInfo(ctx context.Context, session string) (model.SessionInfo, error) {
	var sessionInfo model.SessionInfo

	query := "SELECT * FROM sessions WHERE session = $1"
	err := s.pgx.QueryRow(ctx, query, session).Scan(&sessionInfo.UserID, &sessionInfo.SessionID, &sessionInfo.CreatedAt)
	return sessionInfo, err
}

func (s *sessions) Delete(ctx context.Context, session string) error {
	query := "DELETE FROM sessions WHERE session = $1"
	_, err := s.pgx.Exec(ctx, query, session)
	return err
}
