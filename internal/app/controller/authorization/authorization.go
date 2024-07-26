package authorization

import (
	"StoryPlatforn_GIN/internal/app/service"
	"StoryPlatforn_GIN/internal/domain/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type AuthController struct {
	auth service.Authorization
}

func NewAuthController(auth service.Authorization) AuthController {
	return AuthController{
		auth: auth,
	}
}

func (a AuthController) SignUp(c *gin.Context) {
	var input model.Input

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": model.ErrValidationInput.Error()})
		return
	}

	user, err := a.auth.SignUp(c, input)
	if err != nil {
		if errors.Is(err, model.ErrUsernameTaken) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (a AuthController) SignIn(c *gin.Context) {
	var input model.Input

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": model.ErrValidationInput.Error()})
		return
	}

	user, err := a.auth.SignIn(c, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"session": user.Session})
}

func (a AuthController) Logout(c *gin.Context) {
	session := c.GetHeader("session")
	if session == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": model.SessionEmpty.Error()})
		return
	}

	err := a.auth.Logout(c, session)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
}
