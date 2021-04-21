// routes/routes.go
package routes

import (
	"gohiringchannels/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	//Login
	r.POST("/login", controllers.Login)
	r.POST("/todo", controllers.CreateTodo)
	r.POST("/refresh", controllers.Refresh)
	r.POST("/logout", controllers.Logout)
	//Engineers Endpoint
	r.GET("/api/v1/engineers/getAll")
	//Core Endpoint
	// r.GET("/tasks", controllers.FindTasks)
	// r.POST("/tasks", controllers.CreateTask)
	// r.GET("/tasks/:id", controllers.FindTask)
	// r.PATCH("/tasks/:id", controllers.UpdateTask)
	// r.DELETE("tasks/:id", controllers.DeleteTask)
	return r
}
