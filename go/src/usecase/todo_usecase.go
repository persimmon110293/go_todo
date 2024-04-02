package usecase

import (
	"encoding/json"
	entity "main/domain/entity"
	"main/repository"
)

type ITodoUsecase interface {
	GetAllTodo() ([]entity.Todo, error)
	GetTodoById(string) (*entity.Todo, error)
	CreateTodo([]byte) (*entity.Todo, error)
	UpdateTodoById(string, []byte) (*entity.Todo, error)
	DeleteTodoById(string) error
}

var message map[string]string

type TodoUsecase struct {
	repository repository.ITodoRepository
}

func NewTodoUsecase(repository repository.ITodoRepository) ITodoUsecase {
	return &TodoUsecase{repository}
}

func (u *TodoUsecase) GetAllTodo() ([]entity.Todo, error) {
	todos, err := u.repository.GetAllTodo()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (u *TodoUsecase) GetTodoById(id string) (*entity.Todo, error) {
	todo, err := u.repository.GetTodoById(id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (u *TodoUsecase) CreateTodo(body []byte) (*entity.Todo, error) {
	// JSONをパース
	if err := json.Unmarshal(body, &message); err != nil {
		return nil, err
	}

	return u.repository.CreateTodo(message)
}

func (u *TodoUsecase) UpdateTodoById(id string, body []byte) (*entity.Todo, error) {
	// JSONをパース
	if err := json.Unmarshal(body, &message); err != nil {
		return nil, err
	}

	return u.repository.UpdateTodoById(id, message)
}

func (u *TodoUsecase) DeleteTodoById(id string) error {
	return u.repository.DeleteTodoById(id)
}
