package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sanyathecreator/task-manager/internal/model"
	"github.com/sanyathecreator/task-manager/internal/repository"
)

func getTasks(context *gin.Context) {
	tasks, err := repository.GetTasks()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get data from database."})
		return
	}

	context.JSON(http.StatusOK, tasks)
}

func createTask(context *gin.Context) {
	var task model.Task

	err := context.ShouldBind(&task)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	task.CreatedAt = time.Now()
	err = repository.SaveTask(&task)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save task. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Task created!", "task": task})
}

func updateTask(context *gin.Context) {
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse task id."})
		return
	}

	_, err = repository.GetTaskById(taskId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch task. Try again later."})
		return
	}

	var task model.Task

	err = context.ShouldBind(&task)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	task.ID = taskId
	err = repository.UpdateTask(task)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update task. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Task updated", "task": task})
}

func deleteTask(context *gin.Context) {
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse task id."})
		return
	}

	task, err := repository.GetTaskById(taskId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch task. Try again later."})
		return
	}

	err = repository.DeleteTask(*task)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete task."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Task deleted succesfully!"})
}
