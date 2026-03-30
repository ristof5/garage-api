package routes

import (
	"github.com/gin-gonic/gin"
	"garage-api/controllers"
)

func SetupRoutes(r *gin.Engine, vc *controllers.VehicleController) {

	r.GET("/vehicles", vc.GetVehicles)

	r.POST("/vehicles", vc.CreateVehicle)

	r.PUT("/vehicles/:id", vc.UpdateVehicle)

	r.DELETE("/vehicles/:id", vc.DeleteVehicle)

}