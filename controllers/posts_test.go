package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPosts(t *testing.T) {
	router := SetupRouter(PostController{})
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/", nil)

	fmt.Println("===============>", w.Body.String())
	require.NoError(t, err)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)

}
