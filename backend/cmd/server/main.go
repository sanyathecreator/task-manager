package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanyathecreator/task-manager/internal/db"
	"github.com/sanyathecreator/task-manager/internal/handler"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		context.Header("Access-Control-Allow-Headers", "Content-Type")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.Next()
	})

	handler.RegisterRoutes(server)

	server.Run(":8080")
}
