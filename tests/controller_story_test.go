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

func (s *APITestSuite) TestCreateStoryController() {
	gin.SetMode(gin.TestMode)

	testTable := []struct {
		name        string
		userID      string
		input       model.StoryInput
		expectedErr error
	}{
		{"Correct data", "2", model.StoryInput{Title: "testData", Text: "12345678910111213141516"}, nil},
		{"Long title", "2", model.StoryInput{Title: "ttttttttttttttttttttttttttttttttttttttttt", Text: "12345678910111213141516"}, model.ErrValidationInput},
		{"Empty title", "2", model.StoryInput{Title: "", Text: "12345678910111213141516"}, model.ErrValidationInput},
		{"Empty text", "2", model.StoryInput{Title: "testData", Text: ""}, model.ErrValidationInput},
	}

	for _, test := range testTable {
		s.T().Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			body, err := json.Marshal(test.input)
			if err != nil {
				s.T().Fatal(err)
			}

			req := httptest.NewRequest("POST", "/create", bytes.NewBuffer(body))
			c.Request = req
			c.Set("userID", test.userID)

			s.controllers.Story.CreateStory(c)
			switch test.expectedErr {
			case nil:
				assert.Equal(s.T(), 200, w.Code)
			case model.ErrValidationInput:
				assert.Equal(s.T(), 400, w.Code)
			default:
				s.T().Fatal("unexpected error")
			}
		})
	}
}

func (s *APITestSuite) TestGetStoryController() {
	gin.SetMode(gin.TestMode)

	testTable := []struct {
		name        string
		id          string
		expectedErr error
	}{
		{"Correct data", "2", nil},
		{"Empty id", "", model.IDEmpty},
		{"No data", "0", model.ErrNoData},
	}

	for _, test := range testTable {
		s.T().Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest("GET", "/story", nil)
			c.Params = gin.Params{gin.Param{Key: "id", Value: test.id}}
			c.Request = req

			s.controllers.Story.GetStory(c)
			switch test.expectedErr {
			case nil:
				assert.Equal(s.T(), 200, w.Code)
			case model.IDEmpty:
				assert.Equal(s.T(), 400, w.Code)
			case model.ErrNoData:
				assert.Equal(s.T(), 400, w.Code)
			default:
				s.T().Fatal("unexpected error")
			}
		})
	}
}

func (s *APITestSuite) TestUpdateStoryController() {
	gin.SetMode(gin.TestMode)

	testTable := []struct {
		name        string
		id          string
		userID      string
		input       model.StoryInput
		expectedErr error
	}{
		{"Correct data", "2", "2", model.StoryInput{Title: "testData", Text: "12345678910111213141516"}, nil},
		{"Empty id", "", "1", model.StoryInput{Title: "testData", Text: "12345678910111213141516"}, model.IDEmpty},
		{"Empty userID", "1", "", model.StoryInput{Title: "testData", Text: "12345678910111213141516"}, model.UserIDEmpty},
		{"Long title", "1", "1", model.StoryInput{Title: "ttttttttttttttttttttttttttttttttttttttttt", Text: "12345678910111213141516"}, model.ErrValidationInput},
		{"Empty title", "1", "1", model.StoryInput{Title: "", Text: "12345678910111213141516"}, model.ErrValidationInput},
		{"Empty text", "1", "1", model.StoryInput{Title: "testData", Text: ""}, model.ErrValidationInput},
		{"No data", "0", "2", model.StoryInput{Title: "testData", Text: "12345678910111213141516"}, model.ErrNoData},
	}

	for _, test := range testTable {
		s.T().Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			body, err := json.Marshal(test.input)
			if err != nil {
				s.T().Fatal(err)
			}

			req := httptest.NewRequest("PATCH", "/update", bytes.NewBuffer(body))
			c.Params = gin.Params{gin.Param{Key: "id", Value: test.id}}
			c.Request = req
			c.Set("userID", test.userID)

			s.controllers.Story.UpdateStory(c)
			switch test.expectedErr {
			case nil:
				assert.Equal(s.T(), 200, w.Code)
			case model.IDEmpty:
				assert.Equal(s.T(), 400, w.Code)
			case model.UserIDEmpty:
				assert.Equal(s.T(), 500, w.Code)
			case model.ErrValidationInput:
				assert.Equal(s.T(), 400, w.Code)
			case model.ErrNoData:
				assert.Equal(s.T(), 400, w.Code)
			default:
				s.T().Fatal("unexpected error")
			}
		})
	}
}

func (s *APITestSuite) TestRateStoryController() {
	gin.SetMode(gin.TestMode)

	testTable := []struct {
		name        string
		id          string
		userID      string
		rate        int
		expectedErr error
	}{
		{"Correct data ", "2", "2", -1, nil},
		{"Empty id", "", "1", 1, model.IDEmpty},
		{"Empty userID", "1", "", 1, model.UserIDEmpty},
		{"Rate again 0", "1", "1", 1, model.ErrRateAgain},
		{"Rate again 1", "2", "1", 1, model.ErrRateAgain},
		{"No data", "0", "2", 1, model.ErrNoData},
		{"Validation err (min)", "2", "2", -20, model.ErrValidationInput},
		{"Validation err (0)", "2", "2", 0, model.ErrValidationInput},
		{"Validation err (max)", "2", "2", 20, model.ErrValidationInput},
		//{"No data", "1", "0", 1, model.ErrNoData}, //todo check case with wrong user id
	}

	for _, test := range testTable {
		s.T().Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			body, err := json.Marshal(model.Rate{Rating: test.rate})
			if err != nil {
				s.T().Fatal(err)
			}

			req := httptest.NewRequest("PATCH", "/rate", bytes.NewBuffer(body))
			c.Params = gin.Params{gin.Param{Key: "id", Value: test.id}}
			c.Request = req
			c.Set("userID", test.userID)

			s.controllers.Story.RateStory(c)
			switch test.expectedErr {
			case nil:
				assert.Equal(s.T(), 200, w.Code)
			case model.IDEmpty:
				assert.Equal(s.T(), 400, w.Code)
			case model.UserIDEmpty:
				assert.Equal(s.T(), 500, w.Code)
			case model.ErrRateAgain:
				assert.Equal(s.T(), 400, w.Code)
			case model.ErrNoData:
				assert.Equal(s.T(), 400, w.Code)
			case model.ErrValidationInput:
				assert.Equal(s.T(), 400, w.Code)
			default:
				s.T().Fatal("unexpected error")
			}
		})
	}
}

func (s *APITestSuite) TestDeleteStoryController() {
	gin.SetMode(gin.TestMode)

	testTable := []struct {
		name        string
		id          string
		userID      string
		expectedErr error
	}{
		{"Correct data", "1", "1", nil},
		{"Empty id", "", "2", model.IDEmpty},
		{"Empty userID", "2", "", model.UserIDEmpty},
		{"No data 0", "0", "2", model.ErrNoData},
		{"No data 1", "2", "0", model.ErrNoData},
	}

	for _, test := range testTable {
		s.T().Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest("DELETE", "/delete", nil)
			c.Params = gin.Params{gin.Param{Key: "id", Value: test.id}}
			c.Request = req
			c.Set("userID", test.userID)

			s.controllers.Story.DeleteStory(c)
			switch test.expectedErr {
			case nil:
				assert.Equal(s.T(), 200, w.Code)
			case model.IDEmpty:
				assert.Equal(s.T(), 400, w.Code)
			case model.UserIDEmpty:
				assert.Equal(s.T(), 500, w.Code)
			case model.ErrNoData:
				assert.Equal(s.T(), 400, w.Code)
			default:
				s.T().Fatal("unexpected error")
			}
		})
	}
}
