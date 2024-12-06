package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gotask1/models"
)

func assignTask(context *gin.Context) {
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("Error==>>>>>>>", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}
	_, err = models.GetTaskById(taskId)
	if err != nil {
		fmt.Println("Error==>>>>>>>", err)
		context.JSON(http.StatusNotFound, gin.H{"message": "Unable to Find the Task Id"})
		return
	}

	var assign models.TaskAssignment
	err = context.ShouldBindJSON(&assign)
	if err != nil {
		fmt.Println("error===>>>>", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse request data. Try again later"})
		return
	}
	member_id := assign.MemberId
	_, err = models.GetMemberById(member_id)
	if err != nil {
		fmt.Println("Error==>>>>>>>", err)
		context.JSON(http.StatusNotFound, gin.H{"message": "Unable to Find the Member Id"})
		return
	}
	assign.MemberId = member_id
	assign.TaskId = taskId
	assign.CreatedAt = time.Now()
	assign.UpdatedAt = time.Now()

	err = assign.Save()
	if err != nil {
		fmt.Println("error===>>>>", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to Save data"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully Assigned Task!", "assignment task": assign})
}

func unAssignTask(context *gin.Context) {
	assignId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}

	assignment, err := models.GetMemberById(assignId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch assignment"})
		return
	}

	err = assignment.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to Delete assignment"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully Deleted assignment!!"})
}

func getAssignmentByTaskId(context *gin.Context) {
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}

	assignment, err := models.GetAssignmentByTaskId(taskId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch task assignment of this task Id"})
		return
	}
	context.JSON(http.StatusOK, assignment)
}

func getAssignmentByMemberId(context *gin.Context) {
	memberId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}

	assignment, err := models.GetAssignmentByMemberId(memberId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch task assignment of this member Id"})
		return
	}
	context.JSON(http.StatusOK, assignment)
}
func getAllAssignments(context *gin.Context) {
	assignments, err := models.GetAllAssignments()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to fetch members"})
		return
	}
	context.JSON(http.StatusOK, assignments)
}
