package authorization

import (
	"StoryPlatforn_GIN/internal/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a AuthController) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := c.GetHeader("session")
		if session == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": model.SessionEmpty.Error()})
			return
		}

		info, err := a.auth.GetSessionInfo(c, session)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": model.ErrWrongSession.Error()})
			return
		}

		c.Set("userID", info.UserID)
		c.Next()
	}
}
