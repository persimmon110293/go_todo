package handler

import (
	"main/repository"
	"main/usecase"

	"github.com/gin-gonic/gin"
)

func GetAllTodo(c *gin.Context) {
	repository := repository.NewTodoRepository()
	usecase := usecase.NewTodoUsecase(repository)

	test := usecase.GetAllTodo()

	c.JSON(200, gin.H{
		"message": test,
	})
}
