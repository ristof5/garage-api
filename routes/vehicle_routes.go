package routes

import (
	"garage-api/controllers"
	"garage-api/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	vc *controllers.VehicleController,
	sc *controllers.ServiceController,
	ac *controllers.AuthController,
) {
	// CORS — izinkan Flutter akses backend
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))

	// Public
	r.POST("/auth/login", ac.Login)

	// Protected — semua role yang sudah login
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())

	auth.GET("/vehicles", vc.GetVehicles)
	auth.GET("/services", sc.GetServices)
	auth.GET("/vehicles/:id/services", sc.GetServicesByVehicle)

	// Admin only
	admin := auth.Group("/")
	admin.Use(middlewares.RoleMiddleware("admin"))

	admin.POST("/vehicles", vc.CreateVehicle)
	admin.PUT("/vehicles/:id", vc.UpdateVehicle)
	admin.DELETE("/vehicles/:id", vc.DeleteVehicle)
	admin.POST("/services", sc.CreateService)
}