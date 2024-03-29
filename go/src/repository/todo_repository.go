package repository

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ITodoRepository interface {
	GetAllTodo() ([]Todo, error)
	GetTodoById(string) (*Todo, error)
	CreateTodo(map[string]string) (*Todo, error)
}

type TodoRepository struct{}

func NewTodoRepository() ITodoRepository {
	return &TodoRepository{}
}

func (r *TodoRepository) GetAllTodo() ([]Todo, error) {
	db, err := dbInit()
	if err != nil {
		return nil, err
	}

	todo := []Todo{}
	err = db.Find(&todo).Error
	return todo, err
}

func (r *TodoRepository) GetTodoById(id string) (*Todo, error) {
	db, err := dbInit()
	if err != nil {
		return nil, err
	}

	todo := Todo{}
	err = db.First(&todo, "id = ?", id).Error
	return &todo, err
}

func (r *TodoRepository) CreateTodo(message map[string]string) (*Todo, error) {
	db, err := dbInit()
	if err != nil {
		return nil, err
	}

	todo := Todo{
		Title:       message["Title"],
		Description: message["Description"],
	}
	result := db.Create(&todo)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("failed to create todo")
	}

	return &todo, nil
}

func dbInit() (*gorm.DB, error) {
	dsn := "root:password@tcp(mysql:3306)/practice?charset=utf8mb4&parseTime=true" // 練習用のためハードコーディング
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
