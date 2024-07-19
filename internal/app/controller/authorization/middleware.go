package authorization

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a AuthController) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := c.GetHeader("session")
		if session == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Session header is missing"})
			return
		}

		info, err := a.auth.GetSessionInfo(c, session)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong session"})
			return
		}

		//sessionExpireTime := info.CreatedAt.Add(60 * time.Hour) //todo do fix it
		//if sessionExpireTime.Before(time.Now()) {
		//	err := a.auth.Logout(c, session)
		//	if err != nil {
		//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		//		return
		//	}
		//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Session time is expired"})
		//	return
		//}

		c.Set("userID", info.UserID)
		c.Next()
	}
}
