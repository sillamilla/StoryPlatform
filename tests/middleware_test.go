package tests

import (
	"StoryPlatforn_GIN/internal/domain/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func (s *APITestSuite) TestAuthMiddleware() {
	gin.SetMode(gin.TestMode)

	testTable := []struct {
		name        string
		session     string
		expectedErr error
	}{
		{"Correct session", "session2", nil},
		{"Empty session", "", model.SessionEmpty},
		{"Wrong session", "wrongSession", model.ErrWrongSession},
	}

	for _, test := range testTable {
		s.T().Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request = httptest.NewRequest("POST", "/signup", nil)
			c.Request.Header.Set("session", test.session)

			s.controllers.Authorization.AuthMiddleware()(c)
			switch test.expectedErr {
			case nil:
				assert.Equal(s.T(), 200, w.Code)
			case model.SessionEmpty:
				assert.Equal(s.T(), 400, w.Code)
			case model.ErrWrongSession:
				assert.Equal(s.T(), 400, w.Code)
			default:
				s.T().Fatal("unexpected error")
			}
		})
	}
}
