package usecase

import (
	mock "main/mock/repository"
	"main/repository"
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
	mockRepository.EXPECT().GetAllTodo().Return([]repository.Todo{}, nil)

	// usecaseをインスタンス化
	usecase := NewTodoUsecase(mockRepository)
	todo, err := usecase.GetAllTodo()

	assert.Equal(t, todo, []repository.Todo{})
	assert.Nil(t, err)
}
