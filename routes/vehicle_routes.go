package routes

import (
	"garage-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, vc *controllers.VehicleController, sc *controllers.ServiceController, ac *controllers.AuthController) {
	// vehicle
	r.GET("/vehicles", vc.GetVehicles)

	r.GET("/vehicles/:id", vc.GetVehicleById)

	r.POST("/vehicles", vc.CreateVehicle)

	r.PUT("/vehicles/:id", vc.UpdateVehicle)

	r.DELETE("/vehicles/:id", vc.DeleteVehicle)

	r.GET("/vehicles/:id/detail", vc.GetVehicleDetail)

	// services
	r.POST("/services", sc.CreateService)
	r.GET("/services", sc.GetServices)

	// relation
	r.GET("/vehicles/:id/services", sc.GetServicesByVehicle)

	// auth
	r.POST("/auth/login", ac.Login)

}
