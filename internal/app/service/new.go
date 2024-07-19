package service

import (
	"StoryPlatforn_GIN/internal/app/repository"
)

type Service struct {
	Story         Story
	User          User
	Session       Session
	Authorization Authorization
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Story:         NewStory(repo.Story, repo.User),
		User:          NewUser(repo.User),
		Session:       NewSession(repo.Session),
		Authorization: NewAuthorization(NewSession(repo.Session), NewUser(repo.User)),
	}
}
