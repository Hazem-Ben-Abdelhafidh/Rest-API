package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"rest-api/mocks"
	"rest-api/models"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type PostsSuite struct {
	suite.Suite
	dbMocker             PostController
	postgresProviderMock *mocks.PostService
	testingServer        *httptest.Server
}

func TestPostsSuite(t *testing.T) {
	suite.Run(t, &PostsSuite{})
}

func (ps *PostsSuite) SetupSuite() {
	// Initialize the mock
	postgresProviderMock := new(mocks.PostService)

	// Initialize the controller with the mock service
	dbMocker := NewPostController(postgresProviderMock)

	// Setup the router with the controller
	router := SetupRouter(dbMocker)

	// Create the testing server
	testingServer := httptest.NewServer(router)

	// Set the fields in the suite
	ps.dbMocker = dbMocker
	ps.postgresProviderMock = postgresProviderMock
	ps.testingServer = testingServer

}

func (ps *PostsSuite) TearDownSuite() {
	defer ps.testingServer.Close()
}

func (ps *PostsSuite) TestGetPosts() {
	mockPosts := []models.Post{

		{
			Id:          1,
			Title:       "hazem's first mocked post",
			Description: "hazem's first mocked description",
		},
		{
			Id:          2,
			Title:       "hazem's second mocked post",
			Description: "hazem's second mocked description",
		},
	}

	ps.postgresProviderMock.On("GetPosts").Return(mockPosts, nil)

	response, err := http.Get(fmt.Sprintf("%s/", ps.testingServer.URL))
	ps.Require().NoError(err)
	defer response.Body.Close()
	responseBody := ResponseBody{}
	err = responseBody.FromJson(response.Body)
	ps.Require().NoError(err)
	ps.Require().Equal(responseBody.Status, "success")
	data, err := json.Marshal(responseBody.Data)
	ps.Require().NoError(err)
	mockPostsMarshalled, err := json.Marshal(mockPosts)
	ps.Require().NoError(err)
	ps.Require().JSONEq(string(data), string(mockPostsMarshalled))
	ps.Require().Equal(http.StatusOK, response.StatusCode)

	ps.postgresProviderMock.AssertExpectations(ps.T())
}

func (ps *PostsSuite) TestGetPost() {

	mockPost := models.Post{
		Id:          1,
		Title:       "hazem's first mocked post",
		Description: "hazem's first mocked description",
	}
	tests := []struct {
		name          string
		postId        int
		mock          func(postId int)
		checkResponse func(status int, body io.ReadCloser)
	}{
		{
			name:   "OK",
			postId: 1,
			mock: func(postId int) {
				ps.postgresProviderMock.On("GetPostById", postId).Return(mockPost, nil)

			},
			checkResponse: func(status int, body io.ReadCloser) {
				ps.Require().Equal(status, http.StatusOK)
				responseBody := ResponseBody{}
				err := responseBody.FromJson(body)
				ps.Require().NoError(err)
				ps.Require().Equal(responseBody.Status, "success")
				data, err := json.Marshal(responseBody.Data)
				ps.Require().NoError(err)
				mockPostMarshalled, err := json.Marshal(mockPost)
				ps.Require().NoError(err)
				ps.Require().JSONEq(string(data), string(mockPostMarshalled))
			},
		}, {
			name:   "Invalid Id",
			postId: 0,
			mock: func(postId int) {
				ps.postgresProviderMock.On("GetPostById", postId).Return(models.Post{}, gorm.ErrRecordNotFound)
			},
			checkResponse: func(status int, body io.ReadCloser) {
				ps.Require().Equal(status, http.StatusInternalServerError)
				errorBody := ErrorBody{}
				err := errorBody.FromJson(body)
				ps.Require().NoError(err)
				ps.Require().Equal(errorBody.Status, "error")
				ps.Require().Equal(errorBody.Message, gorm.ErrRecordNotFound.Error())
			},
		},

		{
			name:   "NOT FOUND",
			postId: 69,
			mock: func(postId int) {
				ps.postgresProviderMock.On("GetPostById", postId).Return(models.Post{}, gorm.ErrRecordNotFound)
			},
			checkResponse: func(status int, body io.ReadCloser) {
				ps.Require().Equal(status, http.StatusInternalServerError)
				errorBody := ErrorBody{}
				err := errorBody.FromJson(body)
				ps.Require().NoError(err)
				ps.Require().Equal(errorBody.Status, "error")
				ps.Require().Equal(errorBody.Message, gorm.ErrRecordNotFound.Error())
			},
		},
	}

	for _, test := range tests {
		ps.Run(test.name, func() {
			ps.T().Log("Running test : ", test.name)
			test.mock(test.postId)
			response, err := http.Get(fmt.Sprintf("%s/%d", ps.testingServer.URL, test.postId))
			ps.Require().NoError(err)
			defer response.Body.Close()
			test.checkResponse(response.StatusCode, response.Body)

		})

	}

	ps.postgresProviderMock.AssertExpectations(ps.T())
}

func (ps *PostsSuite) TestDeletePost() {
	tests := []struct {
		name          string
		postId        int
		mock          func(postId int)
		checkResponse func(status int, body io.ReadCloser)
	}{
		{
			name:   "OK",
			postId: 1,
			mock: func(postId int) {
				ps.postgresProviderMock.On("DeletePost", postId).Return(nil)

			},
			checkResponse: func(status int, body io.ReadCloser) {
				ps.Require().Equal(status, http.StatusOK)
				responseBody := ResponseBody{}
				err := responseBody.FromJson(body)
				ps.Require().NoError(err)
				ps.Require().Equal(responseBody.Status, "success")
			},
		}, {
			name:   "Invalid Id",
			postId: 0,
			mock: func(postId int) {
				ps.postgresProviderMock.On("DeletePost", postId).Return(gorm.ErrRecordNotFound)
			},
			checkResponse: func(status int, body io.ReadCloser) {
				ps.Require().Equal(status, http.StatusInternalServerError)
				errorBody := ErrorBody{}
				err := errorBody.FromJson(body)
				ps.Require().NoError(err)
				ps.Require().Equal(errorBody.Status, "error")
				ps.Require().Equal(errorBody.Message, gorm.ErrRecordNotFound.Error())
			},
		},
	}

	for _, test := range tests {
		ps.Run(test.name, func() {
			ps.T().Log("Running test : ", test.name)
			test.mock(test.postId)
			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%d", ps.testingServer.URL, test.postId), nil)
			ps.Require().NoError(err)
			response, err := ps.testingServer.Client().Do(req)
			ps.Require().NoError(err)
			defer response.Body.Close()
			test.checkResponse(response.StatusCode, response.Body)

		})

	}
	ps.postgresProviderMock.AssertExpectations(ps.T())

}

func (ps *PostsSuite) TestCreatePost() {
	mockPost := models.Post{
		Id:          1,
		Title:       "post ok",
		Description: "post ok description",
	}
	tests := []struct {
		name          string
		post          string
		mock          func(post models.PostPayload)
		checkResponse func(status int, body io.ReadCloser)
	}{
		{
			name: "OK",
			post: `{"title": "post ok", "description": "post ok description"}`,
			mock: func(post models.PostPayload) {
				ps.postgresProviderMock.On("CreatePost", post).Return(mockPost, nil)

			},
			checkResponse: func(status int, body io.ReadCloser) {
				ps.Require().Equal(http.StatusCreated, status)
				responseBody := ResponseBody{}
				err := responseBody.FromJson(body)
				ps.Require().NoError(err)
				ps.Require().Equal(responseBody.Status, "success")
				data, err := json.Marshal(responseBody.Data)
				ps.Require().NoError(err)
				mockPostMarshalled, err := json.Marshal(mockPost)
				ps.Require().NoError(err)
				ps.Require().JSONEq(string(data), string(mockPostMarshalled))

			},
		}, {
			name: "Bad Request",
			post: `{"title": "hazem", "description": 1, "crazy": 111}`,
			mock: func(post models.PostPayload) {
				ps.postgresProviderMock.On("CreatePost", post).Return(models.Post{}, mock.Anything)
			},
			checkResponse: func(status int, body io.ReadCloser) {
				ps.Require().Equal(status, http.StatusBadRequest)
				errorBody := ErrorBody{}
				err := errorBody.FromJson(body)
				ps.Require().NoError(err)
				ps.Require().Equal(errorBody.Status, "error")
			},
		},
	}

	for _, test := range tests {
		ps.Run(test.name, func() {
			ps.T().Log("Running test : ", test.name)
			test.mock(models.PostPayload{
				Title:       "post ok",
				Description: "post ok description",
			})
			req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/", ps.testingServer.URL), bytes.NewBuffer([]byte(test.post)))
			ps.Require().NoError(err)
			response, err := ps.testingServer.Client().Do(req)
			ps.Require().NoError(err)
			defer response.Body.Close()
			test.checkResponse(response.StatusCode, response.Body)

		})

	}
	ps.postgresProviderMock.AssertExpectations(ps.T())
}

func (ps *PostsSuite) TestUpdatePost() {
	mockPost := models.Post{
		Id:          1,
		Title:       "post ok",
		Description: "post ok description",
	}
	tests := []struct {
		name          string
		postId        int
		post          string
		mock          func(post models.PostPayload, postId uint)
		checkResponse func(status int, body io.ReadCloser)
	}{
		{
			name:   "OK",
			postId: 1,
			post:   `{"title": "post ok", "description": "post ok description"}`,
			mock: func(post models.PostPayload, postId uint) {
				ps.postgresProviderMock.On("UpdatePost", post, postId).Return(mockPost, nil)

			},
			checkResponse: func(status int, body io.ReadCloser) {
				ps.Require().Equal(http.StatusOK, status)
				responseBody := ResponseBody{}
				err := responseBody.FromJson(body)
				ps.Require().NoError(err)
				ps.Require().Equal(responseBody.Status, "success")
				data, err := json.Marshal(responseBody.Data)
				ps.Require().NoError(err)
				mockPostMarshalled, err := json.Marshal(mockPost)
				ps.Require().NoError(err)
				ps.Require().JSONEq(string(data), string(mockPostMarshalled))
			},
		}, {
			name:   "Bad Request",
			postId: 1,
			post:   `{"title": "hazem", "description": 1, "crazy": 111}`,
			mock: func(post models.PostPayload, postId uint) {
				ps.postgresProviderMock.On("UpdatePost", post, postId).Return(models.Post{}, mock.Anything)
			},
			checkResponse: func(status int, body io.ReadCloser) {
				ps.Require().Equal(status, http.StatusBadRequest)
				errorBody := ErrorBody{}
				err := errorBody.FromJson(body)
				ps.Require().NoError(err)
				ps.Require().Equal(errorBody.Status, "error")
			},
		},
		{
			name:   "Invalid Id",
			postId: 0,
			post:   `{"title": "hazem", "description": 1, "crazy": 111}`,
			mock: func(post models.PostPayload, postId uint) {
				ps.postgresProviderMock.On("UpdatePost", post, postId).Return(models.Post{}, mock.Anything)
			},
			checkResponse: func(status int, body io.ReadCloser) {
				ps.Require().Equal(status, http.StatusBadRequest)
				errorBody := ErrorBody{}
				err := errorBody.FromJson(body)
				ps.Require().NoError(err)
				ps.Require().Equal(errorBody.Status, "error")
			},
		},
	}

	for _, test := range tests {
		ps.Run(test.name, func() {
			ps.T().Log("Running test : ", test.name)
			test.mock(models.PostPayload{
				Title:       "post ok",
				Description: "post ok description",
			}, uint(test.postId))
			req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%d", ps.testingServer.URL, test.postId), bytes.NewBuffer([]byte(test.post)))
			ps.Require().NoError(err)
			response, err := ps.testingServer.Client().Do(req)
			ps.Require().NoError(err)
			defer response.Body.Close()
			test.checkResponse(response.StatusCode, response.Body)
		})

	}
	ps.postgresProviderMock.AssertExpectations(ps.T())
}
