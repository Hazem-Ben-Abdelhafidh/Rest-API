package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(postController PostController) *gin.Engine {
	router := gin.Default()
	router.GET("/", postController.GetPosts)
	router.POST("/", postController.CreatePost)
	router.GET("/:id", postController.GetPost)
	router.DELETE("/:id", postController.DeletePost)
	router.PATCH("/:id", postController.UpdatePost)
	return router
}
