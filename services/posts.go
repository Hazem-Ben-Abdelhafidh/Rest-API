package services

import (
	"rest-api/interfaces"
	"rest-api/models"
)

type PostService struct {
	PostRepo interfaces.PostRepository
}

func NewPostService(postRepo interfaces.PostRepository) PostService {
	return PostService{
		PostRepo: postRepo,
	}
}

func (ps PostService) GetPostById(id int) (models.Post, error) {
	return ps.PostRepo.GetPostById(id)
}

func (ps PostService) GetPosts() ([]models.Post, error) {
	return ps.PostRepo.GetPosts()
}

func (ps PostService) CreatePost(post models.PostPayload) (models.Post, error) {
	return ps.PostRepo.CreatePost(post)
}

func (ps PostService) DeletePost(id int) error {
	return ps.PostRepo.DeletePost(id)
}

func (ps PostService) UpdatePost(post models.PostPayload, postId uint) (models.Post, error) {
	postToCreate := models.Post{
		Id:          postId,
		Title:       post.Title,
		Description: post.Description,
	}
	return ps.PostRepo.UpdatePost(postToCreate)
}
