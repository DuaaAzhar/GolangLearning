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
	server.GET("projects/:id/tasks")
	server.POST("projects/:id/tasks")
	server.GET("tasks/:id")
	server.PUT("tasks/:id")
	server.DELETE("tasks/:id")

	// Members routes
	server.GET("team-members/")
	server.POST("team-members/")
	server.GET("team-members/:id")
	server.DELETE("team-members/:id")

	// Task Assignment Routes
	server.PUT("/tasks/{id}/assign")

}
