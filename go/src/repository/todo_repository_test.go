package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// FIXME: テストコード書く
// DNSでエラーが発生する
func TestGetAllTodo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repository := NewTodoRepository()

		todos, err := repository.GetAllTodo()

		assert.NotNil(t, todos)
		assert.Nil(t, err)
	})
}
