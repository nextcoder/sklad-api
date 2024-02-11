package routes

import (
	"main/src/controllers"
	"main/src/middlewares"

	"github.com/gin-gonic/gin"
)

func startupsGroupRouter(baseRouter *gin.RouterGroup) {
	startups := baseRouter.Group("/auth")

	startups.POST("/login", controllers.Login)
	startups.POST("/register", controllers.Register)
}

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	versionRouter := r.Group("/api")

	startupsGroupRouter(versionRouter)

	versionRouter.GET("/products",
		middlewares.AuthMiddleware(), controllers.GetProducts)
	versionRouter.GET("/products/:id",
		middlewares.AuthMiddleware(), controllers.GetOneProduct)
	versionRouter.POST("/products",
		middlewares.AuthMiddleware(), controllers.CreateProduct)
	versionRouter.DELETE("/products/:id",
		middlewares.AuthMiddleware(), controllers.DeleteProduct)
	versionRouter.PUT("/products/:id",
		middlewares.AuthMiddleware(), controllers.UpdateProduct)
	// versionRouter.POST("/products/move",
	// 	middlewares.AuthMiddleware(), controllers.CreateMovement)
	return r
}
