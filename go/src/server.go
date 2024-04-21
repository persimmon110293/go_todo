package main

import (
	"main/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pong",
		})
	})

	router.GET("/todo", handler.GetAllTodo)
	router.POST("/todo", handler.CreateTodo)
	router.GET("/todo/:id", handler.GetTodoById)
	router.PUT("/todo/:id", handler.UpdateTodoById)
	router.DELETE("/todo/:id", handler.DeleteTodoById)

	router.Run(":8080")
}
