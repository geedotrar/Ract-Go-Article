package main

import (
	"article-api/config"
	"article-api/controllers"
	"article-api/models"
	repositories "article-api/repostiories"
	"article-api/routes"
	"article-api/services"
)

func main() {
	config.Connect()
	config.DB.AutoMigrate(&models.Post{})

	postRepo := repositories.NewPostRepository(config.DB)
	postService := services.NewPostService(postRepo)
	postController := controllers.NewPostController(postService)

	router := routes.SetupRouter(postController)

	router.Run(":8080")
}
