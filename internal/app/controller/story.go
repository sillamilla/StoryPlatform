package Controllers

import (
	"StoryPlatforn_GIN/internal/app/service"
	"StoryPlatforn_GIN/internal/domain/model"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StoryController struct {
	Story service.Story
}

func NewStory(srv service.Story) StoryController {
	return StoryController{
		Story: srv,
	}
}

type error string

func (e error) Error() string {
	return string(e)
}

func (s *StoryController) CreateStory(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": model.UserIDEmpty.Error()})
		return
	}

	var input model.StoryInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := s.Story.CreateStory(c, userID.(string), input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": data})
}

func (s *StoryController) GetStory(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": model.IDEmpty.Error()})
		return
	}

	data, err := s.Story.GetStory(c, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": model.ErrNoData.Error()})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"result": data})
}

func (s *StoryController) UpdateStory(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": model.IDEmpty.Error()})
		return
	}

	userID, ok := c.Get("userID")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": model.UserIDEmpty.Error()})
		return
	}

	var input model.StoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.Story.UpdateStory(c, userID.(string), id, input)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": model.ErrNoData.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (s *StoryController) RateStory(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": model.IDEmpty.Error()})
		return
	}

	userID, ok := c.Get("userID")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": model.UserIDEmpty.Error()})
		return
	}

	var input model.Rate

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.Story.RateStory(c, userID.(string), id, input.Rating)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": model.ErrNoData.Error()})
			return
		}
		if errors.Is(err, model.ErrRateAgain) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": model.ErrRateAgain.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (s *StoryController) DeleteStory(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": model.IDEmpty.Error()})
		return
	}

	userID, ok := c.Get("userID")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": model.UserIDEmpty.Error()})
		return
	}

	err := s.Story.DeleteStory(c, userID.(string), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": model.ErrNoData.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
