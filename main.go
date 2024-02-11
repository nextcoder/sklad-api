package main

import (
	"main/src/models"
	"main/src/routes"
	"main/src/utils"
)

type AuthDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	utils.LoadEnv()

	models.OpenDatabaseConnection()
	models.AutoMigrateModels()
	router := routes.SetupRoutes()

	router.Run(":3000")
}
