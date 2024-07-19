package gin

import (
	"StoryPlatforn_GIN/internal/app/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(controller Controllers.Controller) *gin.Engine {
	router := gin.Default()

	// Public routes
	setupPublicRoutes(router, controller)

	// Routes requiring authentication
	authGroup := router.Group("").Use(controller.Authorization.AuthMiddleware())
	setupAuthRoutes(authGroup, controller)

	return router
}

func setupPublicRoutes(router gin.IRoutes, controller Controllers.Controller) {
	router.POST("/signIn", controller.Authorization.SignIn)
	router.POST("/signUp", controller.Authorization.SignUp)
	router.GET("/story/:id", controller.Story.GetStory)
}

func setupAuthRoutes(group gin.IRoutes, controller Controllers.Controller) {
	group.DELETE("/logout", controller.Authorization.Logout)
	group.POST("/create", controller.Story.CreateStory)
	group.PATCH("/rate/:id", controller.Story.RateStory)

	group.PATCH("/update/:id", controller.Story.UpdateStory)
	group.DELETE("/delete/:id", controller.Story.DeleteStory)
}
