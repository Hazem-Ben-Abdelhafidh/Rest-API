package main

import (
	"fmt"
	"rest-api/controllers"
	"rest-api/db"
	_ "rest-api/docs"
	"rest-api/models"
	"rest-api/repositories"
	"rest-api/services"
)

// @title	Posts REST API
func main() {
	db := db.InitDb()
	err := db.AutoMigrate(&models.Post{})
	if err != nil {
		fmt.Println("error while migrating :", err)
	}

	postRepo := repositories.NewPostGresDb(db)
	postService := services.NewPostService(postRepo)
	postController := controllers.NewPostController(postService)
	router := controllers.SetupRouter(postController)
	router.Run()

}
