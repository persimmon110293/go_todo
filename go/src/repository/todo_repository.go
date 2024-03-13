package repository

type ITodoRepository interface {
	GetAllTodo() string
}

type TodoRepository struct{}

func NewTodoRepository() ITodoRepository {
	return &TodoRepository{}
}

func (r *TodoRepository) GetAllTodo() string {
	// TODO: 実装する
	return "GetAllTodo"
}
