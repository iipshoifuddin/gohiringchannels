package main

import (
	"gohiringchannels/config"
	"gohiringchannels/models"
	"gohiringchannels/routes"
	"os"
)

func main() {
	config.ENVSetup()
	db := config.SetupDB()
	db.AutoMigrate(&models.Engineer{})
	db.AutoMigrate(&models.Enterprice{})

	r := routes.SetupRoutes(db)
	r.Run(os.Getenv("PORT_RUNNING")) //setting Port

}
