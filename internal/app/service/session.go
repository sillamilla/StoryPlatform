package service

import (
	"StoryPlatforn_GIN/internal/app/repository"
	"StoryPlatforn_GIN/internal/domain/model"
	"context"
	"github.com/pkg/errors"
	time2 "time"
)

type Session interface {
	GetUserID(ctx context.Context, session string) (string, error)
	CreateOrUpdate(ctx context.Context, userID, session string) error
	GetSessionInfo(ctx context.Context, session string) (model.SessionInfo, error)
	Logout(ctx context.Context, session string) error
}

type session struct {
	repo repository.Session
}

func NewSession(repo repository.Session) Session {
	return session{repo: repo}
}

func (s session) GetUserID(ctx context.Context, session string) (string, error) {
	const op = "session.GetUserID"

	id, err := s.repo.GetUserID(ctx, session)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	return id, nil
}
func (s session) GetSessionInfo(ctx context.Context, session string) (model.SessionInfo, error) {
	const op = "session.GetSessionInfo"

	info, err := s.repo.GetSessionInfo(ctx, session)
	if err != nil {
		return model.SessionInfo{}, errors.Wrap(err, op)
	}

	return info, nil
}

func (s session) CreateOrUpdate(ctx context.Context, userID, session string) error {
	const op = "session.CreateOrUpdate"
	//todo time2 wtf
	time := time2.Now()
	err := s.repo.Upsert(ctx, userID, session, time)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (s session) Logout(ctx context.Context, session string) error {
	const op = "session.Logout"

	err := s.repo.Delete(ctx, session)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}
