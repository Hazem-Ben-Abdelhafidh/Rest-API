package repositories

import (
	"rest-api/models"

	"gorm.io/gorm"
)

type PostgresDb struct {
	Db *gorm.DB
}

func NewPostGresDb(db *gorm.DB) PostgresDb {
	return PostgresDb{
		Db: db,
	}
}

func (db PostgresDb) GetPostById(id int) (models.Post, error) {
	var post models.Post
	err := db.Db.Table("posts").First(&post, id).Error
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func (db PostgresDb) CreatePost(payload models.PostPayload) (models.Post, error) {
	post := models.Post{
		Title:       payload.Title,
		Description: payload.Description,
	}
	err := db.Db.Table("posts").Create(&post).Error
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func (db PostgresDb) GetPosts() ([]models.Post, error) {
	var posts []models.Post
	err := db.Db.Table("posts").Find(&posts).Error
	if err != nil {
		return []models.Post{}, err
	}
	return posts, nil
}

func (db PostgresDb) DeletePost(id int) error {
	return db.Db.Table("posts").Delete(&models.Post{}, id).Error
}

func (db PostgresDb) UpdatePost(post models.Post) (models.Post, error) {
	err := db.Db.Table("posts").Save(&post).Error
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}
