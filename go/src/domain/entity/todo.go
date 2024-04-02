package entity

type Todo struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}
