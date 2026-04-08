package routes

import (
	"garage-api/controllers"
	"garage-api/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, vc *controllers.VehicleController, sc *controllers.ServiceController, ac *controllers.AuthController) {

	// public route
	r.POST("/auth/login", ac.Login)

	// protected routes
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())

	// vehicles & services user
	auth.GET("/vehicles", vc.GetVehicles)
	auth.GET("/services", sc.GetServices)
	auth.GET("/vehicles/:id/services", sc.GetServicesByVehicle)

	// admin role
	admin := auth.Group("/")
	admin.Use(middlewares.RoleMiddleware("admin"))

	admin.POST("/vehicles", vc.CreateVehicle)
	admin.PUT("/vehicles/:id", vc.UpdateVehicle)
	admin.DELETE("/vehicles/:id", vc.DeleteVehicle)

	admin.POST("/services", sc.CreateService)

}
