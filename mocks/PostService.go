// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	models "rest-api/models"

	mock "github.com/stretchr/testify/mock"
)

// PostService is an autogenerated mock type for the PostService type
type PostService struct {
	mock.Mock
}

// CreatePost provides a mock function with given fields: post
func (_m *PostService) CreatePost(post models.PostPayload) (models.Post, error) {
	ret := _m.Called(post)

	if len(ret) == 0 {
		panic("no return value specified for CreatePost")
	}

	var r0 models.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(models.PostPayload) (models.Post, error)); ok {
		return rf(post)
	}
	if rf, ok := ret.Get(0).(func(models.PostPayload) models.Post); ok {
		r0 = rf(post)
	} else {
		r0 = ret.Get(0).(models.Post)
	}

	if rf, ok := ret.Get(1).(func(models.PostPayload) error); ok {
		r1 = rf(post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePost provides a mock function with given fields: id
func (_m *PostService) DeletePost(id int) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeletePost")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPostById provides a mock function with given fields: id
func (_m *PostService) GetPostById(id int) (models.Post, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetPostById")
	}

	var r0 models.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (models.Post, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) models.Post); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Post)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPosts provides a mock function with given fields:
func (_m *PostService) GetPosts() ([]models.Post, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPosts")
	}

	var r0 []models.Post
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Post, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Post); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Post)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePost provides a mock function with given fields: post, postId
func (_m *PostService) UpdatePost(post models.PostPayload, postId uint) (models.Post, error) {
	ret := _m.Called(post, postId)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePost")
	}

	var r0 models.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(models.PostPayload, uint) (models.Post, error)); ok {
		return rf(post, postId)
	}
	if rf, ok := ret.Get(0).(func(models.PostPayload, uint) models.Post); ok {
		r0 = rf(post, postId)
	} else {
		r0 = ret.Get(0).(models.Post)
	}

	if rf, ok := ret.Get(1).(func(models.PostPayload, uint) error); ok {
		r1 = rf(post, postId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPostService creates a new instance of PostService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPostService(t interface {
	mock.TestingT
	Cleanup(func())
}) *PostService {
	mock := &PostService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
