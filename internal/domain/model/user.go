package model

import "time"

type Input struct {
	Username string `form:"username" json:"username" binding:"required,min=4,max=20"`
	Password string `form:"password" json:"password" binding:"required,min=7,max=30"`
}

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Session   string    `json:"session"`
	CreatedAt time.Time `json:"createdAt"`
}

func UserFromInput(ID string, session string, user Input) User {
	return User{
		ID:       ID,
		Username: user.Username,
		Password: user.Password,
		Session:  session,
	}
}
