package repository

import (
	"errors"
	entity "main/domain/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ITodoRepository interface {
	GetAllTodo() ([]entity.Todo, error)
	GetTodoById(string) (*entity.Todo, error)
	CreateTodo(map[string]string) (*entity.Todo, error)
	UpdateTodoById(string, map[string]string) (*entity.Todo, error)
	DeleteTodoById(string) error
}

type TodoRepository struct{}

func NewTodoRepository() ITodoRepository {
	return &TodoRepository{}
}

func (r *TodoRepository) GetAllTodo() ([]entity.Todo, error) {
	db, err := dbInit()
	if err != nil {
		return nil, err
	}

	todo := []entity.Todo{}
	err = db.Find(&todo).Error
	return todo, err
}

func (r *TodoRepository) GetTodoById(id string) (*entity.Todo, error) {
	db, err := dbInit()
	if err != nil {
		return nil, err
	}

	todo := entity.Todo{}
	err = db.First(&todo, "id = ?", id).Error
	return &todo, err
}

func (r *TodoRepository) CreateTodo(message map[string]string) (*entity.Todo, error) {
	db, err := dbInit()
	if err != nil {
		return nil, err
	}

	todo := entity.Todo{
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

func (r *TodoRepository) UpdateTodoById(id string, message map[string]string) (*entity.Todo, error) {
	db, err := dbInit()
	if err != nil {
		return nil, err
	}

	todo := entity.Todo{
		Title:       message["Title"],
		Description: message["Description"],
	}
	result := db.Model(&entity.Todo{}).Where("id = ?", id).Updates(todo)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("failed to update todo")
	}

	return &todo, nil
}

func (r *TodoRepository) DeleteTodoById(id string) error {
	db, err := dbInit()
	if err != nil {
		return err
	}

	result := db.Delete(&entity.Todo{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to delete todo")
	}

	return nil
}

func dbInit() (*gorm.DB, error) {
	dsn := "root:password@tcp(mysql:3306)/practice?charset=utf8mb4&parseTime=true" // 練習用のためハードコーディング
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
