package tests

import (
	"StoryPlatforn_GIN/internal/domain/model"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func (s *APITestSuite) TestSignUpController() {
	gin.SetMode(gin.TestMode)

	testTable := []struct {
		name        string
		input       model.Input
		expectedErr error
	}{
		{"Correct data", model.Input{Username: "testData", Password: "testData"}, nil},
		{"Username taken", model.Input{Username: "username1", Password: "testData"}, model.ErrUsernameTaken},
		{"Short username", model.Input{Username: "t", Password: "testData"}, model.ErrValidationInput},
		{"Short password", model.Input{Username: "testData1", Password: "t"}, model.ErrValidationInput},
		{"Long password", model.Input{Username: "testData1", Password: "tttttttttttttttttttttttttttttttt"}, model.ErrValidationInput},
		{"Long username", model.Input{Username: "tttttttttttttttttttttttttttttttt", Password: "testData1"}, model.ErrValidationInput},
	}

	for _, test := range testTable {
		s.T().Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			body, err := json.Marshal(test.input)
			if err != nil {
				s.T().Fatal(err)
			}

			req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(body))
			c.Request = req

			s.controllers.Authorization.SignUp(c)
			switch test.expectedErr {
			case nil:
				assert.Equal(s.T(), 201, w.Code)
			case model.ErrUsernameTaken:
				assert.Equal(s.T(), 400, w.Code)
			case model.ErrValidationInput:
				assert.Equal(s.T(), 400, w.Code)
			default:
				s.T().Fatal("unexpected error")
			}
		})
	}
}

func (s *APITestSuite) TestSignInController() {
	gin.SetMode(gin.TestMode)

	testTable := []struct {
		name        string
		input       model.Input
		expectedErr error
	}{
		{"Correct data", model.Input{Username: "username1", Password: "password1"}, nil},
		{"Incorrect password", model.Input{Username: "username1", Password: "incorrect"}, model.InvalidPassword},
		{"Incorrect username", model.Input{Username: "incorrect", Password: "password1"}, model.ErrUserNotFound},
		{"Short username", model.Input{Username: "t", Password: "testData"}, model.ErrValidationInput},
		{"Short password", model.Input{Username: "username2", Password: "t"}, model.ErrValidationInput},
		{"Long password", model.Input{Username: "username2", Password: "tttttttttttttttttttttttttttttttt"}, model.ErrValidationInput},
		{"Long username", model.Input{Username: "tttttttttttttttttttttttttttttttt", Password: "password2"}, model.ErrValidationInput},
	}

	for _, test := range testTable {
		s.T().Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			body, err := json.Marshal(test.input)
			if err != nil {
				s.T().Fatal(err)
			}

			req := httptest.NewRequest("POST", "/signIn", bytes.NewBuffer(body))
			c.Request = req

			s.controllers.Authorization.SignIn(c)
			switch test.expectedErr {
			case nil:
				assert.Equal(s.T(), 200, w.Code)
			case model.InvalidPassword:
				assert.Equal(s.T(), 400, w.Code)
			case model.ErrUserNotFound:
				assert.Equal(s.T(), 400, w.Code)
			case model.ErrValidationInput:
				assert.Equal(s.T(), 400, w.Code)
			default:
				s.T().Fatal("unexpected error")
			}
		})
	}
}

func (s *APITestSuite) TestLogoutController() {
	gin.SetMode(gin.TestMode)

	testTable := []struct {
		name        string
		session     string
		expectedErr error
	}{
		{"Correct data", "session1", nil},
		{"Empty session", "", model.SessionEmpty},
	}

	for _, test := range testTable {
		s.T().Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest("DELETE", "/logout", nil)
			c.Request = req
			c.Request.Header.Set("session", test.session)

			s.controllers.Authorization.Logout(c)
			switch test.expectedErr {
			case nil:
				assert.Equal(s.T(), 200, w.Code)
			case model.SessionEmpty:
				assert.Equal(s.T(), 400, w.Code)
			default:
				s.T().Fatal("unexpected error")
			}
		})
	}
}
