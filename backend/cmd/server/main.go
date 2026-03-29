package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sanyathecreator/task-manager/internal/db"
	"github.com/sanyathecreator/task-manager/internal/handler"
)

func main() {
	db.InitDB()

	server := gin.Default()

	handler.RegisterRoutes(server)

	server.Run(":8080")
}
