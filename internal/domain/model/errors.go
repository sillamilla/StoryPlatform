package model

type error string

func (e error) Error() string {
	return string(e)
}

const (
	ErrUsernameTaken   error = "this username is taken"
	ErrUserNotFound    error = "user not found"
	ErrNoData          error = "no data found with given name"
	ErrRateAgain       error = "you can not rate again"
	ErrValidationInput error = "validation input error"
	InvalidPassword    error = "invalid password"
	ErrWrongSession    error = "wrong session"

	UserIDEmpty  error = "user id is empty"
	IDEmpty      error = "id is empty"
	SessionEmpty error = "session is empty"
)
