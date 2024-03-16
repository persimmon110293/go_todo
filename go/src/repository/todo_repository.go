package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ITodoRepository interface {
	GetAllTodo() ([]Todo, error)
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

func dbInit() (*gorm.DB, error) {
	dsn := "root:password@tcp(mysql:3306)/practice?charset=utf8mb4&parseTime=true" // 練習用のためハードコーディング
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
