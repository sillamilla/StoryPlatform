package service

import (
	"StoryPlatforn_GIN/internal/app/repository"
	"StoryPlatforn_GIN/internal/domain/model"
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"time"
)

type User interface {
	Create(ctx context.Context, user model.User) error
	IsUsernameAvailable(ctx context.Context, username string) (bool, error)
	GetByUsername(ctx context.Context, username string) (model.User, error)
	GetByID(ctx context.Context, id string) (model.User, error)
	GetUserIdBySession(ctx context.Context, session string) (string, error)
}

type user struct {
	repo repository.User
}

func NewUser(repo repository.User) User {
	return &user{repo: repo}
}

func (u user) Create(ctx context.Context, user model.User) error {
	const op = "user.Create"

	user.CreatedAt = time.Now()
	if err := u.repo.Create(ctx, user); err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (u user) GetByUsername(ctx context.Context, username string) (model.User, error) {
	const op = "user.GetByUsername"

	byUsername, err := u.repo.GetByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, errors.New("No user with such username") //todo not work
		}
		return model.User{}, errors.Wrap(err, op)
	}

	return byUsername, nil
}

func (u user) GetByID(ctx context.Context, id string) (model.User, error) {
	const op = "user.GetByID"

	byID, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return model.User{}, errors.Wrap(err, op)
	}

	return byID, nil
}

func (u user) GetUserIdBySession(ctx context.Context, session string) (string, error) {
	const op = "user.GetBySession"

	id, err := u.repo.GetUserIdBySession(ctx, session)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	return id, nil
}

func (u user) IsUsernameAvailable(ctx context.Context, username string) (bool, error) {
	const op = "user.IsUsernameAvailable"

	isUsernameExist, err := u.repo.IsUsernameExist(ctx, username)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	return isUsernameExist, nil
}
