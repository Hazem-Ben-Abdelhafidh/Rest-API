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

// GetPosts	godoc
// @Summary	Get All Posts
// @Description This endpoint is used to get all posts
// @Produce application/json
// @Success 200 {object} ResponseBody{}
// @Tags posts
// @Router / [get]
func (pc PostController) GetPosts(c *gin.Context) {
	posts, err := pc.postService.GetPosts()
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	respondWithJson(c, http.StatusOK, posts)
}

// GetPost	godoc
// @Summary	Get one Post
// @Description This endpoint is used to get one post by passing it's id
// @Param id path string true "get post by id"
// @Produce application/json
// @Success 200 {object} ResponseBody{}
// @Tags posts
// @Router /{id} [get]
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

// CreatePost	godoc
// @Summary	Create Post
// @Description This endpoint is used to create a new post
// @Param post body models.PostPayload true "create Post"
// @Produce application/json
// @Success 201 {object} ResponseBody{}
// @Tags posts
// @Router / [post]
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

// DeletePost	godoc
// @Summary	Delete a post
// @Description This endpoint is used to delete a post by id
// @Param id path string true "delete post by id"
// @Produce application/json
// @Success 200 {object} ResponseBody{}
// @Tags posts
// @Router /{id} [delete]
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

// UpdatePost	godoc
// @Summary	Update Post
// @Description This endpoint is used to update an existing Post by it's id
// @Param post body models.PostPayload true "updatePost Post"
// @Param id path string true "update post by id"
// @Produce application/json
// @Success 200 {object} ResponseBody{}
// @Tags posts
// @Router /{id} [patch]
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
