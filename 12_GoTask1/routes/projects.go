package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gotask1/models"
)

func getProjects(context *gin.Context) {
	projects, err := models.GetAllProjects()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to fetch projects"})
		return
	}
	context.JSON(http.StatusOK, projects)
}

func createProject(context *gin.Context) {
	var project models.Project
	err := context.ShouldBindJSON(&project)
	if err != nil {
		fmt.Println("error===>>>", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}
	project.ID = 1
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()

	fmt.Println("project ===>>", project)
	err = project.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to save project"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Project created Successfully ", "project": project})
}

func getProjectById(context *gin.Context) {
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
	context.JSON(http.StatusOK, project)
}

func updateProject(context *gin.Context) {
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

func deleteProject(context *gin.Context) {
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
