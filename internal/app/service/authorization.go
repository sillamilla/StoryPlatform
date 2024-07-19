package service

import (
	"StoryPlatforn_GIN/internal/app/service/helper"
	"StoryPlatforn_GIN/internal/domain/model"
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Authorization interface {
	SignUp(ctx context.Context, user model.Input) (model.User, error)
	SignIn(ctx context.Context, req model.Input) (model.User, error)
	Logout(ctx context.Context, session string) error
	GetSessionInfo(ctx context.Context, session string) (model.SessionInfo, error)
}

type auth struct {
	session Session
	user    User
}

func NewAuthorization(sessions Session, user User) Authorization {
	return auth{session: sessions, user: user}
}

func (a auth) SignUp(ctx context.Context, input model.Input) (model.User, error) {
	const op = "authorization.SignUp"

	ok, err := a.user.IsUsernameAvailable(ctx, input.Username)
	if err != nil {
		return model.User{}, errors.Wrap(err, op)
	}

	if ok {
		return model.User{}, errors.New("This username is taken")
	}

	password, err := helper.HashPassword(input.Password)
	if err != nil {
		return model.User{}, errors.Wrap(err, op)
	}
	input.Password = password

	session, err := helper.GenerateSession()
	if err != nil {
		return model.User{}, errors.Wrap(err, op)
	}

	id := uuid.NewString()
	newUser := model.UserFromInput(id, session, input)

	err = a.user.Create(ctx, newUser)
	if err != nil {
		return model.User{}, errors.Wrap(err, op)
	}

	err = a.session.CreateOrUpdate(ctx, id, session)
	if err != nil {
		return model.User{}, errors.Wrap(err, op) //todo to many wrap
	}

	return newUser, nil
}

func (a auth) SignIn(ctx context.Context, input model.Input) (model.User, error) {
	const op = "authorization.SignIn"

	user, err := a.user.GetByUsername(ctx, input.Username)
	if err != nil {
		return model.User{}, errors.Wrap(err, op)
	}

	err = helper.ComparePassword(user.Password, input.Password)
	if err != nil {
		return model.User{}, errors.Wrap(err, "Invalid password")
	}

	if user.Session == "" {
		session, err := helper.GenerateSession()
		if err != nil {
			return model.User{}, errors.Wrap(err, op)
		}
		user.Session = session
	}

	input.Password = user.Password

	err = a.session.CreateOrUpdate(ctx, user.ID, user.Session)
	if err != nil {
		return model.User{}, errors.Wrap(err, op)
	}

	return user, nil
}

func (a auth) GetSessionInfo(ctx context.Context, session string) (model.SessionInfo, error) {
	const op = "authorization.GetSessionInfo"

	info, err := a.session.GetSessionInfo(ctx, session)
	if err != nil {
		return model.SessionInfo{}, errors.Wrap(err, op)
	}

	return info, nil
}

func (a auth) Logout(ctx context.Context, session string) error {
	const op = "authorization.Logout"

	err := a.session.Logout(ctx, session)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}
