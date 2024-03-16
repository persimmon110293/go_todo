package usecase

import (
	"main/repository"
)

type TodoUsecase struct {
	repository repository.ITodoRepository
}

func NewTodoUsecase(repository repository.ITodoRepository) *TodoUsecase {
	return &TodoUsecase{repository}
}

func (u *TodoUsecase) GetAllTodo() ([]repository.Todo, error) {
	todos, err := u.repository.GetAllTodo()
	if err != nil {
		return nil, err
	}

	return todos, nil
}
