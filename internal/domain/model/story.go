package model

import (
	"time"
)

type StoryInput struct {
	Title string `form:"title" json:"title" binding:"required,min=5,max=40"`
	Text  string `form:"text" json:"text" binding:"required,min=15,max=700"`
}

type Story struct {
	ID        string
	UserID    string
	Author    string
	Title     string
	Text      string
	Rating    int
	CreatedAt time.Time
}

type Rate struct {
	Rating int `json:"rating" binding:"required,min=-1,max=1"`
}

func StoryFromInput(id string, userID string, author string, time time.Time, input StoryInput) Story {
	return Story{
		ID:        id,
		UserID:    userID,
		Author:    author,
		Title:     input.Title,
		Text:      input.Text,
		Rating:    0,
		CreatedAt: time,
	}
}
