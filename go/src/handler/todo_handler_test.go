package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// FIXME: テストコード書く
// DNSでエラーが発生する
func TestGetAllTodo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := gin.Default()
		w := httptest.NewRecorder() // httptestの使い方調べてみる

		req, _ := http.NewRequest("GET", "/todo", nil)
		router.GET("/todo", GetAllTodo)
		router.ServeHTTP(w, req)

		assert.NotNil(t, w.Body)
		assert.Equal(t, w.Code, 200)
	})
}
