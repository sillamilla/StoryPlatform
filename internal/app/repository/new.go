package repository

import (
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Story   Story
	User    User
	Session Session
}

func New(db *pgx.Conn) *Repository {
	return &Repository{
		Story:   NewStory(db),
		User:    NewUser(db),
		Session: NewSession(db),
	}
}
