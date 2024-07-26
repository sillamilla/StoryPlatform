package tests

import (
	"StoryPlatforn_GIN/internal/app/service/helper"
	"StoryPlatforn_GIN/internal/domain/model"
	"time"
)

type Rate struct {
	UserID  string
	StoryID string
}

var (
	password1, _ = helper.HashPassword("password1")
	password2, _ = helper.HashPassword("password2")
	users        = []model.User{
		{
			ID:        "1",
			Username:  "username1",
			Password:  password1,
			Session:   "session1",
			CreatedAt: time.Now(),
		},
		{
			ID:        "2",
			Username:  "username2",
			Password:  password2,
			Session:   "session2",
			CreatedAt: time.Now(),
		},
	}

	stories = []model.Story{
		{
			ID:        "1",
			UserID:    users[0].ID,
			Author:    users[0].Username,
			Title:     "Story Title 1",
			Text:      "This is the text of the first story.",
			Rating:    -3,
			CreatedAt: time.Now(),
		},
		{
			ID:        "2",
			UserID:    users[1].ID,
			Author:    users[1].Username,
			Title:     "Story Title 2",
			Text:      "This is the text of the second story.",
			Rating:    4,
			CreatedAt: time.Now(),
		},
	}

	sessions = []model.SessionInfo{
		{
			SessionID: "session1",
			UserID:    users[0].ID,
			CreatedAt: time.Now(),
		},
		{

			SessionID: "session2",
			UserID:    users[1].ID,
			CreatedAt: time.Now(),
		},
	}

	rates = []Rate{
		{
			UserID:  users[0].ID,
			StoryID: stories[1].ID,
		},
		{
			UserID:  users[1].ID,
			StoryID: stories[0].ID,
		},
	}
)
