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

func TestGetTodoById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepository := mock.NewMockITodoRepository(mockCtrl)
	mockRepository.EXPECT().GetTodoById("1").Return(&entity.Todo{
		Title:       "test title",
		Description: "test description",
	}, nil)

	t.Run("success", func(t *testing.T) {
		usecase := NewTodoUsecase(mockRepository)
		todo, err := usecase.GetTodoById("1")

		assert.Equal(t, todo, &entity.Todo{
			Title:       "test title",
			Description: "test description",
		})
		assert.Nil(t, err)
	})
}

func TestCreateTodo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepository := mock.NewMockITodoRepository(mockCtrl)
	mockRepository.EXPECT().CreateTodo(gomock.Any()).Return(&entity.Todo{
		Title:       "test title",
		Description: "test description",
	}, nil)

	t.Run("success", func(t *testing.T) {
		usecase := NewTodoUsecase(mockRepository)
		todo, err := usecase.CreateTodo([]byte(`{"title": "test title", "description": "test description"}`))

		assert.Equal(t, todo, &entity.Todo{
			Title:       "test title",
			Description: "test description",
		})
		assert.Nil(t, err)
	})
}

func TestUpdateTodoById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMessage := map[string]string{
		"Title":       "mock title",
		"Description": "mock description",
	}

	mockRepository := mock.NewMockITodoRepository(mockCtrl)
	mockRepository.EXPECT().UpdateTodoById("1", mockMessage).Return(&entity.Todo{
		Title:       "mock title",
		Description: "mock description",
	}, nil)

	t.Run("success", func(t *testing.T) {
		usecase := NewTodoUsecase(mockRepository)
		todo, err := usecase.UpdateTodoById("1", []byte(`{"Title": "mock title", "Description": "mock description"}`))

		assert.Equal(t, todo, &entity.Todo{
			Title:       "mock title",
			Description: "mock description",
		})
		assert.Nil(t, err)
	})
}

func TestDeleteTodoById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepository := mock.NewMockITodoRepository(mockCtrl)
	mockRepository.EXPECT().DeleteTodoById("1").Return(nil)

	t.Run("success", func(t *testing.T) {
		usecase := NewTodoUsecase(mockRepository)
		result := usecase.DeleteTodoById("1")

		assert.Nil(t, result)
	})
}
