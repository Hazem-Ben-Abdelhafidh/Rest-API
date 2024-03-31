package controllers

import (
	"net/http"
	"rest-api/interfaces"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService interfaces.PostService
}

func NewPostController(ps interfaces.PostService) PostController {
	return PostController{
		postService: ps,
	}
}

func (pc PostController) GetPosts(c *gin.Context) {
	posts, err := pc.postService.GetPosts()
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	respondWithJson(c, http.StatusOK, posts)
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
	respondWithJson(c, http.StatusOK, post)
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
	respondWithJson(c, http.StatusCreated, post)
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
		respondWithError(c, http.StatusInternalServerError, err)
	}
	respondWithJson(c, http.StatusOK, nil)
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
	post, err := pc.postService.UpdatePost(payload, uint(idInt))
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	respondWithJson(c, http.StatusOK, post)
}
