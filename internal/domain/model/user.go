package model

import "time"

type Input struct {
	Username string `form:"username" json:"username" binding:"required,min=4,max=20"`
	Password string `form:"password" json:"password" binding:"required,min=7,max=30"`
}

type User struct {
	ID        string
	Username  string
	Password  string
	Session   string
	CreatedAt time.Time
}

func UserFromInput(ID string, session string, user Input) User {
	return User{
		ID:       ID,
		Username: user.Username,
		Password: user.Password,
		Session:  session,
	}
}
