package controllers

import (
	"net/http"
	"rest-api/models"
	"rest-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService services.PostService
}

func NewPostController(ps services.PostService) PostController {
	return PostController{
		postService: ps,
	}
}

func respondWithError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"status":  "error",
		"message": err.Error(),
	})
}

func respondWithJson(c *gin.Context, statusCode int, dataKey string, data any) {
	c.JSON(statusCode, gin.H{
		"status": "success",
		dataKey:  data,
	})
}

func (pc PostController) GetPosts(c *gin.Context) {
	posts, err := pc.postService.GetPosts()
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	respondWithJson(c, http.StatusOK, "posts", posts)
}

func (pc PostController) GetPost(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, err)
		return
	}
	post, err := pc.postService.GetPostById(idInt)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	respondWithJson(c, http.StatusOK, "post", post)
}

func (pc PostController) CreatePost(c *gin.Context) {
	var payload models.PostPayload
	if err := c.BindJSON(&payload); err != nil {
		respondWithError(c, http.StatusBadRequest, err)
		return
	}
	post, err := pc.postService.CreatePost(payload)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	respondWithJson(c, http.StatusCreated, "post", post)
}

func (pc PostController) DeletePost(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, err)
		return
	}
	err = pc.postService.DeletePost(idInt)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, err)
	}
	respondWithJson(c, http.StatusOK, "post", nil)
}

func (pc PostController) UpdatePost(c *gin.Context) {
	var payload models.PostPayload
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, err)
		return
	}

	if err = c.BindJSON(&payload); err != nil {
		respondWithError(c, http.StatusBadRequest, err)
		return
	}
	postPayload := models.Post{
		Id:          uint(idInt),
		Title:       payload.Title,
		Description: payload.Description,
	}
	post, err := pc.postService.UpdatePost(postPayload)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	respondWithJson(c, http.StatusOK, "post", post)
}
