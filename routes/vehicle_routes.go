package routes

import (
	"garage-api/controllers"
	"github.com/gin-gonic/gin"
	"garage-api/middlewares"
)

func SetupRoutes(r *gin.Engine, vc *controllers.VehicleController, sc *controllers.ServiceController, ac *controllers.AuthController) {

	// public route
	r.POST("/auth/login", ac.Login)

	// protected routes
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())

	// vehicles
	auth.GET("/vehicles", vc.GetVehicles)
	auth.POST("/vehicles", vc.CreateVehicle)
	auth.PUT("/vehicles/:id", vc.UpdateVehicle)
	auth.DELETE("/vehicles/:id", vc.DeleteVehicle)

	// services
	auth.POST("/services", sc.CreateService)
	auth.GET("/services", sc.GetServices)
	auth.GET("/vehicles/:id/services", sc.GetServicesByVehicle)
}
