package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Register route handlers
	server.GET("/tasks", getTasks)
	server.POST("/tasks", createTask)
	server.PUT("/tasks/:id", updateTask)
	server.DELETE("/tasks/:id", deleteTask)
}
