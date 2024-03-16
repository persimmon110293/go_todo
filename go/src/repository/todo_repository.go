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
	db := dbInit()
	todo := []Todo{}

	err := db.Find(&todo).Error
	return todo, err
}

func dbInit() *gorm.DB {
	dsn := "root:password@tcp(mysql:3306)/practice?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
