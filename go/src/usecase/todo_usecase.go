package usecase

import "main/repository"

type TodoUsecase struct {
	repository repository.ITodoRepository
}

func NewTodoUsecase(repository repository.ITodoRepository) *TodoUsecase {
	return &TodoUsecase{repository}
}

func (u *TodoUsecase) GetAllTodo() string {
	return u.repository.GetAllTodo()
}
