package usecase

import (
	"encoding/json"
	"main/repository"
)

type ITodoUsecase interface {
	GetAllTodo() ([]repository.Todo, error)
	GetTodoById(string) (*repository.Todo, error)
	CreateTodo([]byte) (*repository.Todo, error)
}

type TodoUsecase struct {
	repository repository.ITodoRepository
}

func NewTodoUsecase(repository repository.ITodoRepository) ITodoUsecase {
	return &TodoUsecase{repository}
}

func (u *TodoUsecase) GetAllTodo() ([]repository.Todo, error) {
	todos, err := u.repository.GetAllTodo()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (u *TodoUsecase) GetTodoById(id string) (*repository.Todo, error) {
	todo, err := u.repository.GetTodoById(id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (u *TodoUsecase) CreateTodo(body []byte) (*repository.Todo, error) {
	var message map[string]string

	// JSONをパース
	if err := json.Unmarshal(body, &message); err != nil {
		return nil, err
	}

	return u.repository.CreateTodo(message)
}
