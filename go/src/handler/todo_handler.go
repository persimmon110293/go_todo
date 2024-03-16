package handler

import (
	"main/repository"
	"main/usecase"

	"github.com/gin-gonic/gin"
)

func GetAllTodo(c *gin.Context) {
	// repositoryとusecaseをインスタンス化
	repository := repository.NewTodoRepository()
	usecase := usecase.NewTodoUsecase(repository)

	todos, err := usecase.GetAllTodo()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": todos,
	})
}
