package interfaces

import "rest-api/models"

type PostRepository interface {
	GetPostById(id int) (models.Post, error)
	GetPosts() ([]models.Post, error)
	CreatePost(post models.PostPayload) (models.Post, error)
	DeletePost(id int) error
	UpdatePost(post models.Post) (models.Post, error)
}
