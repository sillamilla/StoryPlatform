package Controllers

import (
	"StoryPlatforn_GIN/internal/app/controller/authorization"
	"StoryPlatforn_GIN/internal/app/service"
)

type Controller struct {
	Story         StoryController
	User          UserController
	Authorization authorization.AuthController
}

func New(srv *service.Service) Controller {
	return Controller{
		Story:         NewStory(srv.Story),
		User:          NewUser(srv.User),
		Authorization: authorization.NewAuthController(srv.Authorization),
	}
}
