package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanyathecreator/task-manager/internal/model"
	"github.com/sanyathecreator/task-manager/internal/repository"
)

func getTasks(context *gin.Context) {

}

func createTask(context *gin.Context) {
	var task model.Task

	err := context.ShouldBind(&task)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = repository.Save(task)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save task. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Task created!", "task": task})
}

func updateTask(context *gin.Context) {

}

func deleteTask(context *gin.Context) {

}
