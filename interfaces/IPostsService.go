package interfaces

import "rest-api/models"

type PostService interface {
	GetPostById(id int) (models.Post, error)
	GetPosts() ([]models.Post, error)
	CreatePost(post models.PostPayload) (models.Post, error)
	DeletePost(id int) error
	UpdatePost(post models.PostPayload, postId uint) (models.Post, error)
}
