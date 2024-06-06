package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTodo(t *testing.T) {
	t.Cleanup(func() {
		os.Unsetenv("DB_HOST")
		os.Unsetenv("ENV_PATH")
	})
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("ENV_PATH", "../.env")

	t.Run("success", func(t *testing.T) {
		repository := NewTodoRepository()

		todos, err := repository.GetAllTodo()

		assert.NotEmpty(t, todos)
		assert.NoError(t, err)
	})
}
