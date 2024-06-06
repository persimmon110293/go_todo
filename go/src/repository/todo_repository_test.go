package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// FIXME: テストコード書く
// DNSでエラーが発生する
func TestGetAllTodo(t *testing.T) {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("ENV_PATH", "../.env")
	t.Run("success", func(t *testing.T) {
		repository := NewTodoRepository()

		todos, err := repository.GetAllTodo()

		assert.NotEmpty(t, todos)
		assert.NoError(t, err)
	})
}
