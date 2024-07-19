package authorization

import (
	"StoryPlatforn_GIN/internal/app/service"
	"StoryPlatforn_GIN/internal/domain/model"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := a.auth.SignUp(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (a AuthController) SignIn(c *gin.Context) {
	var input model.Input

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := a.auth.SignIn(c, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"session": user.Session})
}

func (a AuthController) Logout(c *gin.Context) {
	session := c.GetHeader("session")
	if session == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Session field is empty"})
		return
	}

	err := a.auth.Logout(c, session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
}
