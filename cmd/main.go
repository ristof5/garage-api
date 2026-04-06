package main

import (
	"garage-api/config"
	"garage-api/controllers"
	"garage-api/repositories"
	"garage-api/routes"
	// "garage-api/seeds"
	"garage-api/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.ConnectDB()

	if err != nil {
		panic(err)
	}

	repo := repositories.VehicleRepository{DB: db}
	service := services.VehicleService{Repo: &repo}
	controller := controllers.VehicleController{Service: &service}

	serviceRepo := repositories.ServiceRepository{DB: db}
	serviceService := services.ServiceService{Repo: &serviceRepo}
	serviceController := controllers.ServiceController{Service: &serviceService}

	// user repo
	userRepo := repositories.UserRepository{DB: db}
	// auth service
	authService := services.AuthService{Repo: &userRepo}
	// auth controller
	authController := controllers.AuthController{Service: &authService}
	// seeds.SeedUser(db)

	r := gin.Default()

	routes.SetupRoutes(r, &controller, &serviceController, &authController)

	r.Run(":8080")
}
