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
	projectId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}
	_, err = models.GetProjectById(projectId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Unable to Find the Project Id"})
		return
	}

	tasks, err := models.GetAllTasks(projectId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to fetch tasks"})
		return
	}
	context.JSON(http.StatusOK, tasks)
}

func createTask(context *gin.Context) {
	projectId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}
	_, err = models.GetProjectById(projectId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Unable to Find the Project Id"})
		return
	}

	var task models.Task
	err = context.ShouldBindJSON(&task)
	if err != nil {
		fmt.Println("error===>>>", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}
	task.ID = 1
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	task.Project_ID = projectId
	// validating status
	err = task.SetStatus(task.TaskStatus)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to set Status"})
		return
	}

	fmt.Println("project ===>>", task)
	err = task.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "task created Successfully ", "task": task})
}

func getTaskById(context *gin.Context) {
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}

	task, err := models.GetTaskById(taskId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch task"})
		return
	}
	context.JSON(http.StatusOK, task)
}

func updateTask(context *gin.Context) {
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}

	var task models.Task
	err = context.ShouldBindJSON(&task)
	if err != nil {
		fmt.Println("error===>>>>", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse request data. Try again later"})
		return
	}

	// validating the status
	err = task.SetStatus(task.TaskStatus)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	// validating project id
	_, err = models.GetProjectById(task.Project_ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Project id doesn't exist"})
		return
	}

	//setting taskId
	task.ID = taskId

	err = task.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to Update task"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully updated task!"})

}

func deleteTask(context *gin.Context) {

	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}

	task, err := models.GetTaskById(taskId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch task"})
		return
	}

	_, err = models.GetAssignmentByTaskId(taskId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed! This task has pending assignments"})
		return
	}

	err = task.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to Delete task"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully Deleted task!!"})

}
