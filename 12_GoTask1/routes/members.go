package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gotask1/models"
)

func getMembers(context *gin.Context) {
	members, err := models.GetAllMembers()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to fetch members"})
		return
	}
	context.JSON(http.StatusOK, members)
}

func createMember(context *gin.Context) {
	var member models.Member
	err := context.ShouldBindJSON(&member)
	if err != nil {
		fmt.Println("error===>>>", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}

	member.CreatedAt = time.Now()
	member.UpdatedAt = time.Now()

	fmt.Println("member ===>>", member)
	err = member.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to save Member's Data"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Team Member added Successfully ", "team Member": member})
}

func getMemberById(context *gin.Context) {
	memberId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}

	member, err := models.GetMemberById(memberId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch member"})
		return
	}
	context.JSON(http.StatusOK, member)
}

func updateMember(context *gin.Context) {
	memberId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}

	_, err = models.GetMemberById(memberId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch member"})
		return
	}

	var member models.Member
	err = context.ShouldBindJSON(&member)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse request data. Try again later"})
		return
	}

	member.ID = memberId
	member.UpdatedAt = time.Now()
	err = member.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to Update member"})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully updated member!"})
}

func deleteMember(context *gin.Context) {
	memberId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}

	member, err := models.GetMemberById(memberId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch member"})
		return
	}

	_, err = models.GetAssignmentByMemberId(memberId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed! This Member has pending assignments"})
		return
	}

	err = member.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to Delete member"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully Deleted member!!"})

}
