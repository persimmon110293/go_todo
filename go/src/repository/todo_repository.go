//go:generate mockgen -source=todo_repository.go -destination=../mock/repository/todo_repository_mock.go -package=mock

package repository

import (
	"errors"
	"fmt"
	entity "main/domain/entity"
	"os"

	"github.com/joho/godotenv"
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
	envPath := os.Getenv("ENV_PATH")
	if envPath == "" {
		envPath = "/go/src/app/.env"
	}
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}

	// .envファイルから環境変数を取得してDSNを生成
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
