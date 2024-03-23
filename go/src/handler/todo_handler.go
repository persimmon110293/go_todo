package handler

import (
	"io"
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

func GetTodoById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{
			"message": "id is required",
		})
		return
	}

	repository := repository.NewTodoRepository()
	usecase := usecase.NewTodoUsecase(repository)

	todo, err := usecase.GetTodoById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": todo,
	})
}

func CreateTodo(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	repository := repository.NewTodoRepository()
	usecase := usecase.NewTodoUsecase(repository)

	result, err := usecase.CreateTodo(body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"result":  result,
	})
}
