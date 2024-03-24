package repositories

import (
	"rest-api/db"
	"rest-api/models"
	"rest-api/utils"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type PostsSuite struct {
	suite.Suite
	pg PostgresDb
}

func TestPostsSuite(t *testing.T) {
	suite.Run(t, &PostsSuite{})
}

func (ps *PostsSuite) SetupSuite() {
	db := db.InitDb()
	pg := NewPostGresDb(db)
	ps.pg = pg
	err := ps.pg.Db.Raw("DELETE FROM posts").Error
	ps.Require().NoError(err)
}

func (ps *PostsSuite) AfterTest(suiteTest, testName string) {
	clearPostsTable(ps)
}

func createRandomPost(t *testing.T, pg PostgresDb) models.Post {
	title := utils.RandomString(5)
	description := utils.RandomString(20)
	postToCreate := models.PostPayload{
		Title:       title,
		Description: description,
	}
	createdPost, err := pg.CreatePost(postToCreate)
	require.NoError(t, err)
	require.NotEmpty(t, createdPost)
	require.NotZero(t, createdPost.Id)
	require.Equal(t, createdPost.Title, postToCreate.Title)
	require.Equal(t, createdPost.Description, postToCreate.Description)
	return createdPost
}
func (ps *PostsSuite) TestGetPostById() {
	createdPost := createRandomPost(ps.T(), ps.pg)

	post, err := ps.pg.GetPostById(int(createdPost.Id))
	ps.Require().NoError(err)
	ps.Require().NotEmpty(post)
	ps.Require().Equal(createdPost.Id, post.Id)
	ps.Require().Equal(createdPost.Title, post.Title)
	ps.Require().Equal(createdPost.Description, post.Description)
}

func (ps *PostsSuite) TestCreatePost() {
	createRandomPost(ps.T(), ps.pg)
}

func (ps *PostsSuite) TestGetPosts() {
	for i := 0; i < 10; i++ {
		createRandomPost(ps.T(), ps.pg)
	}
	posts, err := ps.pg.GetPosts()
	ps.Require().NoError(err)
	ps.Require().NotEmpty(posts)
	ps.Require().Equal(len(posts), 10)
}

func (ps *PostsSuite) TestDeletePost() {
	post := createRandomPost(ps.T(), ps.pg)
	err := ps.pg.DeletePost(int(post.Id))
	ps.Require().NoError(err)
	deletedPost, err := ps.pg.GetPostById(int(post.Id))
	ps.Require().Error(err)
	ps.Require().EqualError(err, gorm.ErrRecordNotFound.Error())
	ps.Require().Empty(deletedPost)
}

func (ps *PostsSuite) TestUpdatePost() {
	post := createRandomPost(ps.T(), ps.pg)
	updatedPayload := models.Post{
		Id:          post.Id,
		Title:       "modifed",
		Description: post.Description,
	}
	updatedPost, err := ps.pg.UpdatePost(updatedPayload)
	ps.Require().NoError(err)
	ps.Require().NotEmpty(updatedPost)
	ps.Require().Equal(post.Id, updatedPost.Id)
	ps.Require().Equal(post.Description, updatedPost.Description)
	ps.Require().Equal(updatedPost.Title, updatedPayload.Title)
}

func clearPostsTable(ps *PostsSuite) {
	err := ps.pg.Db.Exec("DELETE FROM posts").Error
	if err != nil {
		ps.T().Log("ERROR WHILE DELETING!!!", err.Error())
	}
	ps.Require().NoError(err)
}
