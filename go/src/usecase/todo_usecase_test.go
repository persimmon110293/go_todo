package usecase

import (
	entity "main/domain/entity"
	mock "main/mock/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTodo(t *testing.T) {
	// mockを管理するコントローラーを作成
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// mockのrepositoryを作成
	mockRepository := mock.NewMockITodoRepository(mockCtrl)
	// 対象のmethodをmock化(ここではGetAllTodo)
	mockRepository.EXPECT().GetAllTodo().Return([]entity.Todo{}, nil)

	t.Run("success", func(t *testing.T) {
		// usecaseをインスタンス化
		usecase := NewTodoUsecase(mockRepository)
		todo, err := usecase.GetAllTodo()

		assert.Equal(t, todo, []entity.Todo{})
		assert.Nil(t, err)
	})
}
