package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	// Project routes
	server.GET("projects", getProjects)
	server.POST("projects", createProject)
	server.GET("projects/:id", getProjectById)
	server.PUT("projects/:id", updateProject)
	server.DELETE("projects/:id", deleteProject)

	// Tasks routes
	server.GET("projects/:id/tasks", getTasks)
	server.POST("projects/:id/tasks", createTask)
	server.GET("tasks/:id", getTaskById)
	server.PUT("tasks/:id", updateTask)
	server.DELETE("tasks/:id", deleteTask)

	// Members routes
	server.GET("team-members/", getMembers)
	server.POST("team-members/", createMember)
	server.GET("team-members/:id", getMemberById)
	server.PUT("team-members/:id", updateMember)
	server.DELETE("team-members/:id", deleteMember)

	// Task Assignment Routes

	server.POST("/tasks/:id/assign", assignTask)
	server.DELETE("/tasks/:id/assign", unAssignTask)
	server.GET("/tasks/:id/assign", getAssignmentByTaskId)
	server.GET("/members/:id/assign", getAssignmentByMemberId)
	server.GET("/assignments", getAllAssignments)

}
