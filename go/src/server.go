package main

import (
	"main/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
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
