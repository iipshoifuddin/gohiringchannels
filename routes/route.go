// routes/routes.go
package routes

import (
	"gohiringchannels/controllers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.Static("/public/engineers", os.Getenv("PATH_STATIC_ENG"))

	//Use RegisterValidation in the Controllers struct entity
	// and use the key are the registered word
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("general", bookableGeneral)
		v.RegisterValidation("bookabledate", bookableDate)
		v.RegisterValidation("passwd", bookablePassword)
		v.RegisterValidation("email", validateEmail)
		v.RegisterValidation("usernm", bookableUsername)
	}

	//Login
	r.POST("/login", controllers.Login)
	r.POST("/todo", controllers.CreateTodo)
	r.POST("/refresh", controllers.Refresh)
	r.POST("/logout", controllers.Logout)
	r.POST("/api/v1/engineers/register", controllers.CreateEngineer)
	//Engineers Endpoint
	r.GET("/api/v1/engineers/getAll", controllers.FineEngineers)
	r.GET("/api/v1/engineers/find/:id", controllers.FineEngineer)
	r.POST("/api/v1/engineers/update/:id", controllers.UpdateEngginer)
	r.POST("/api/v1/engineers/uploadimage/:id", controllers.UploadShowcaseEngginer)

	//Core Endpoint
	return r
}
