package Controllers

import (
	"StoryPlatforn_GIN/internal/app/service"
)

type UserController struct {
	User service.User
}

func NewUser(srv service.User) UserController {
	return UserController{
		User: srv,
	}
}

//type error string
//
//func (e error) Error() string {
//	return string(e)
//}
//
//const ()
