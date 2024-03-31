package controllers

import (
	"encoding/json"
	"io"

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

type ResponseBody struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

func (rb *ResponseBody) ToJson(w io.Writer) error {
	return json.NewEncoder(w).Encode(rb)
}

func (rb *ResponseBody) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(rb)
}

type ErrorBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (eb *ErrorBody) ToJson(w io.Writer) error {
	return json.NewEncoder(w).Encode(eb)
}

func (eb *ErrorBody) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(eb)
}

func respondWithError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"status":  "error",
		"message": err.Error(),
	})
}

func respondWithJson(c *gin.Context, statusCode int, data any) {
	responseBody := ResponseBody{
		Status: "success",
		Data:   data,
	}
	c.JSON(statusCode, responseBody)
}
