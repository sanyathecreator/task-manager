package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sanyathecreator/task-manager/internal/handler"
)

func main() {
	server := gin.Default()

	handler.RegisterRoutes(server)

	server.Run(":8080")
}
