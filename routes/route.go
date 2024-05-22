package routes

import (
	"article-api/controllers"
	middleware "article-api/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter(postController *controllers.PostController) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	api := router.Group("/article")
	{
		api.POST("/posts", postController.CreatePost)
		api.GET("/posts/:id", postController.GetPostByID)
		api.PUT("/posts/:id", postController.UpdatePost)
		api.DELETE("/posts/:id", postController.DeletePost)
		api.GET("/posts", postController.GetAllPosts)
	}

	return router
}
