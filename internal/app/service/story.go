package service

import (
	"StoryPlatforn_GIN/internal/app/repository"
	"StoryPlatforn_GIN/internal/domain/model"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"time"
)

type Story interface {
	CreateStory(ctx context.Context, userID string, input model.StoryInput) (model.Story, error)
	GetStory(ctx context.Context, id string) (model.Story, error)
	RateStory(ctx context.Context, userID string, id string, rate int) error
	UpdateStory(ctx context.Context, userID string, id string, input model.StoryInput) error
	DeleteStory(ctx context.Context, userID string, id string) error
}

type story struct {
	st repository.Story
	us repository.User
}

func NewStory(st repository.Story, us repository.User) Story {
	return &story{st: st, us: us}
}

func (s *story) CreateStory(ctx context.Context, userID string, input model.StoryInput) (model.Story, error) {
	const op = "story.CreateStory"

	user, err := s.us.GetByID(ctx, userID)
	if err != nil {
		return model.Story{}, errors.Wrap(err, op)
	}

	id := uuid.NewString()
	now := time.Now()

	data := model.StoryFromInput(id, user.ID, user.Username, now, input)
	data.CreatedAt = time.Now()

	err = s.st.Create(ctx, data)
	if err != nil {
		return model.Story{}, errors.Wrap(err, op)
	}

	return data, nil
}

func (s *story) GetStory(ctx context.Context, id string) (model.Story, error) {
	const op = "story.GetStory"
	data, err := s.st.Get(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Story{}, model.ErrNoData
		}
		return model.Story{}, errors.Wrap(err, op)
	}

	return data, nil
}

func (s *story) RateStory(ctx context.Context, userID string, id string, rate int) error {
	const op = "story.RateStory"

	data, err := s.st.IsRated(ctx, userID, id)
	if len(data) > 0 {
		return model.ErrRateAgain
	}
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = s.st.MarkUserRated(ctx, userID, id)
			if err != nil {
				return errors.Wrap(err, op)
			}

			err = s.st.Rate(ctx, id, rate)
			if err != nil {
				return errors.Wrap(err, op)
			}
		} else {
			return errors.Wrap(err, op)
		}
	}

	return nil
}

func (s *story) UpdateStory(ctx context.Context, userID string, id string, input model.StoryInput) error {
	const op = "story.UpdateStory"

	err := s.st.Update(ctx, userID, id, input)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (s *story) DeleteStory(ctx context.Context, userID string, id string) error {
	const op = "story.DeleteStory"

	err := s.st.Delete(ctx, userID, id)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}
