package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gotask1/models"
)

func getTasks(context *gin.Context) {
	tasks, err := models.GetAllTasks()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to fetch projects"})
		return
	}
	context.JSON(http.StatusOK, tasks)
}

func createTask(context *gin.Context) {
	var task models.Task
	err := context.ShouldBindJSON(&task)
	if err != nil {
		fmt.Println("error===>>>", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}
	task.ID = 1
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	fmt.Println("project ===>>", task)
	err = task.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to save project"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "task created Successfully ", "project": task})
}

func getTaskById(context *gin.Context) {
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}

	task, err := models.GetTaskById(taskId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch project"})
		return
	}
	context.JSON(http.StatusOK, task)
}

func updateTask(context *gin.Context) {
	projectId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}

	_, err = models.GetProjectById(projectId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch project"})
		return
	}

	var project models.Project
	err = context.ShouldBindJSON(&project)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse request data. Try again later"})
	}

	project.ID = projectId
	err = project.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to Update project"})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully updated project!"})

}

func deleteTask(context *gin.Context) {
	projectId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}

	project, err := models.GetProjectById(projectId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch project"})
		return
	}

	err = project.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to Delete project"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully Deleted project!!"})

}
