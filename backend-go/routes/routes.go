package routes

import (
	"net/http"

	"github.com/MUsamaT/task-tracker/controllers"
	"github.com/gin-gonic/gin"
)

//import "fmt"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Task Track API is running",
		})
	})

	auth := server.Group("/api/auth")

	{
		auth.POST("/register", controllers.RegisterUser)
		auth.POST("/login", controllers.Login)
	}

	tasks := server.Group("/api/")

	{
		tasks.POST("/task", controllers.CreateTask)
		tasks.GET("/tasks", controllers.GetAllTasks)
		tasks.GET("/tasks/:id", controllers.GetUserTasks)
		tasks.PUT("/task/:id", controllers.UpddateTask)
		tasks.DELETE("/task/:id", controllers.DeleTask)
	}

}
