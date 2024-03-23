package main

import (
	"main/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("todo", handler.GetAllTodo)
	router.POST("todo", handler.CreateTodo)
	router.GET("todo/:id", handler.GetTodoById)

	router.Run(":8080")
}
